package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type transfer_req struct {
    From string
	To string
	Coins string
	JWT string
}

func transfer (w http.ResponseWriter, r *http.Request){


	var p transfer_req

    err := json.NewDecoder(r.Body).Decode(&p)
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}


	from,e1:=strconv.Atoi(p.From)
	to,e2:= strconv.Atoi(p.To)
	coin,e3:=strconv.Atoi(p.Coins)
	if(e1!=nil || e2!=nil || e3!=nil){fmt.Fprint(w, "Illformated request"); return}



	
	
	db, err := sql.Open("sqlite3", "./a.db")
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	defer db.Close()





	a:=p.JWT
// fmt.Println(a)
if(len(a)<4){fmt.Fprint(w,"Invalid JWT!"); return}

x:= strings.Split(a,".")
// sig:= strings.Split(x[2],";")[0]
if (len(x)!=3){fmt.Fprint(w,"Invalid JWT!"); return}
body:= x[0]+"." +x[1]
sig:= x[2]


key := []byte("yoyoyoyoyoyo")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(body))
signature:=base64.StdEncoding.EncodeToString(h.Sum(nil))

if(sig!=signature){
	fmt.Fprint(w,"Invalid JWT!")
	return
} else {

	
	head, err:= base64.StdEncoding.DecodeString(x[1])
	if err!=nil{log.Fatalln(err.Error())}
	// s1:= head
	var s1 Jss
	err= json.Unmarshal([]byte("{"+string(head)+"}"), &s1)
	if err!=nil{log.Fatalln(err.Error())}
	a1,_:=strconv.Atoi(s1.User)
	
	if(a1!=from){fmt.Fprintf(w,"Not Authorised!");return}

}











	sqlStatement := `SELECT coins FROM User WHERE id=$1 OR id= $2;`
	row,err := (db.Query(sqlStatement, from,to))
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	
	responses:=0
	for row.Next(){	responses++}

	if(responses<2){fmt.Fprintf(w, "USERS DONT EXIST :: goto /signup to create users and coins\n"); return}




	ctx:=context.Background()
	tx,err:=db.BeginTx(ctx,nil)
	if(err!=nil){log.Fatal(err); return}

	res,err1:= tx.ExecContext(ctx, "UPDATE User SET  coins = coins - $1 WHERE id = $2 AND coins -$1 >=0; ",coin,from)
	val,err2:= res.RowsAffected()

	if (err1!=nil || err2!=nil ){tx.Rollback(); return}
	if(val==0){fmt.Fprint(w, "Balance not Sufficient!\n");tx.Rollback(); return}
	// fmt.Println(from,to,coin,val)

	

	_,err= tx.ExecContext(ctx, "UPDATE User SET  coins = coins + $1 WHERE id = $2;",0.98*float64(coin),to)
	if err!=nil{tx.Rollback(); log.Fatal(err); return}
	
	err= tx.Commit()
	if err!=nil{log.Fatal(err); return}
	

	newtx := `INSERT INTO Tx(Time, coins, txfrom, txto) VALUES (?, ?,?,?)`

	sqlcodetx, err := db.Prepare(newtx)


	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	_, err = sqlcodetx.Exec(time.Now().String(), coin,from, to)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}


	fmt.Fprintf(w,"Done!\n")



}