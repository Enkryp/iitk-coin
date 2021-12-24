package main

import (
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
)

type red_req struct {
	Roll string
	Id   string
	JWT  string
}

func redeem(w http.ResponseWriter, r *http.Request) {

	var p red_req

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}

	roll, e1 := strconv.Atoi(p.Roll)
	id, e2 := strconv.Atoi(p.Id)
	// coin,e3:=strconv.Atoi(p.Coins)
	if e1 != nil || e2 != nil {
		fmt.Fprint(w, "Illformated request")
		return
	}

	db, err := sql.Open("sqlite3", "./a.db")
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	defer db.Close()

	a := p.JWT
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

		if a1 != roll {
			fmt.Fprintf(w, "Not Authorised!")
			return
		}

	}

	sqlStatement := `SELECT pass FROM User WHERE id=$1;`
	row, _ := (db.Query(sqlStatement, roll))
	hi := 0
	for row.Next() {
		// _= row.Scan(&pass2)
		hi++

	}
	if hi == 0 {
		fmt.Fprintf(w, "You do not EXIST ! xD")
		return
	}

	coin := 0
	reci := 1
	sqlStatement = `SELECT coins,Recipient FROM Redeem WHERE Id=$1;`
	row, _ = (db.Query(sqlStatement, id))
	hi = 0
	for row.Next() {
		_ = row.Scan(&coin, &reci)
		hi++

	}
	if hi == 0 {
		fmt.Fprintf(w, "Item does not Exist !")
		return
	}
	if reci != 0 {
		fmt.Fprintf(w, "Item already redeemed !")
		return
	}

	new := `INSERT INTO Pending(awardId,coins,Recipient) VALUES (?, ?, ?)`

	sqlcode, err := db.Prepare(new)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(id, coin, roll)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintf(w, "Request Made to Gensec....\n")

}
