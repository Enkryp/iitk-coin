package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"fmt"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type transfer_req struct {
    From string
	To string
	Coins string
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








	sqlStatement := `SELECT coins FROM User WHERE id=$1 OR id= $2;`
	row,err := (db.Query(sqlStatement, from,to))
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	
	responses:=0
	for row.Next(){	responses++}

	if(responses<2){fmt.Fprintf(w, "USERS DONT EXIST :: goto /create to create users and coins\n"); return}




	ctx:=context.Background()
	tx,err:=db.BeginTx(ctx,nil)
	if(err!=nil){log.Fatal(err); return}

	res,err1:= tx.ExecContext(ctx, "UPDATE User SET  coins = coins - $1 WHERE id = $2 AND coins -$1 >=0; ",coin,from)
	val,err2:= res.RowsAffected()

	if (err1!=nil || err2!=nil ){tx.Rollback(); return}
	if(val==0){fmt.Fprint(w, "Balance not Sufficient!\n");tx.Rollback(); return}
	// fmt.Println(from,to,coin,val)


	_,err= tx.ExecContext(ctx, "UPDATE User SET  coins = coins + $1 WHERE id = $2;",coin,to)
	if err!=nil{tx.Rollback(); log.Fatal(err); return}
	
	err= tx.Commit()
	if err!=nil{log.Fatal(err); return}
	
	fmt.Fprintf(w,"Done!\n")



}