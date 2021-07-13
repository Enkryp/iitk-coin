package main

import (
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

	_ "github.com/mattn/go-sqlite3"
)

type add_req struct {
	Roll  string
	Award string
	Coins string
	JWT   string
}

func add(w http.ResponseWriter, r *http.Request) {

	var p add_req

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatalln(err.Error())

	}

	roll, err := strconv.Atoi(p.Roll)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	coin, err := strconv.Atoi(p.Coins)
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

	sqlStatement := `SELECT coins FROM User WHERE id=$1;`
	row, _ := (db.Query(sqlStatement, roll))

	responses := 0
	for row.Next() {
		responses++
	}

	if responses == 0 {

		fmt.Fprint(w, "User doesnt exist")
		return

	} else {

		new := `INSERT INTO Redeem(Item,coins,Recipient) VALUES ( ?, ?,?)`

		sqlcode, err := db.Prepare(new)

		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = sqlcode.Exec(p.Award, coin, 0)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}

	fmt.Fprintf(w, "Done!\n")

}
