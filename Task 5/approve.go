package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type approve_req struct {
	Roll string
	Id   string
	JWT  string
}

func approve(w http.ResponseWriter, r *http.Request) {

	var p approve_req

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatalln(err.Error())

	}

	roll, err := strconv.Atoi(p.Roll)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	id, err := strconv.Atoi(p.Id)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	JWT := p.JWT
	db, err := sql.Open("sqlite3", "./a.db")
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	defer db.Close()

	a := JWT
	// fmt.Println(a)
	if len(a) < 4 {
		fmt.Fprint(w, "Invalid JWT!")
		return
	}

	x := strings.Split(a, ".")
	// sig:= strings.Split(x[2],";")[0]
	if len(x) != 3 {
		fmt.Fprint(w, "Invalid JWT!")
		return
	}
	body := x[0] + "." + x[1]
	sig := x[2]

	// goo:=0
	key := []byte("yoyoyoyoyoyo")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(body))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if sig != signature {
		fmt.Fprint(w, "Invalid JWT!")
		return
	} else {

		head, err := base64.StdEncoding.DecodeString(x[1])
		if err != nil {
			log.Fatalln(err.Error())
		}
		// s1:= head
		var s1 Jss
		err = json.Unmarshal([]byte("{"+string(head)+"}"), &s1)
		if err != nil {
			log.Fatalln(err.Error())
		}
		a1, _ := strconv.Atoi(s1.User)
		// goo=a1
		st := 1

		// fmt.Println(a1)
		sqlStatement := `SELECT adm FROM User WHERE id=$1;`
		row, _ := (db.Query(sqlStatement, a1))
		hi := 0
		for row.Next() {
			_ = row.Scan(&st)
			hi++

		}
		// fmt.Println(st)
		if hi == 0 || st == 0 {
			fmt.Fprint(w, "Sorry, Not Authorised!")
			return
		}

	}

	awardId := 0
	coin := 0
	reci := 0
	sqlStatement := `SELECT awardId,coins,Recipient FROM Pending WHERE id=$1;`
	row, _ := (db.Query(sqlStatement, id))
	hi := 0
	for row.Next() {
		_ = row.Scan(&awardId, &coin, &reci)
		hi++

	}
	if hi == 0 {
		fmt.Fprintf(w, "Txn doest exist !")
		return
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err1 := tx.ExecContext(ctx, "UPDATE User SET  coins = coins - $1 WHERE id = $2 AND coins -$1 >=0; ", coin, reci)
	val, err2 := res.RowsAffected()

	if err1 != nil || err2 != nil {
		tx.Rollback()
		return
	}
	if val == 0 {
		fmt.Fprint(w, "Balance not Sufficient!\n")
		tx.Rollback()

		new := `DELETE FROM Pending WHERE id= $1`

		sqlcode, err := db.Prepare(new)

		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = sqlcode.Exec(id)
		if err != nil {
			log.Fatalln(err.Error())
		}

		return
	}
	// fmt.Println(from,to,coin,val)

	_, err = tx.ExecContext(ctx, "UPDATE Redeem SET  Recipient = $1 WHERE Id = $2;", roll, awardId)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return
	}

	newtx := `INSERT INTO Tx(Time, coins, txfrom, txto) VALUES (?, ?,?,?)`

	sqlcodetx, err := db.Prepare(newtx)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	_, err = sqlcodetx.Exec(time.Now().String(), coin, "Award:"+strconv.Itoa(id), roll)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	new := `DELETE FROM Pending WHERE id= $1`

	sqlcode, err := db.Prepare(new)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintf(w, "Done!\n")

}
