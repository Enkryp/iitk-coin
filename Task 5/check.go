package main

import (
	"database/sql"
	"encoding/json"
	"strconv"

	// "encoding/json"
	"fmt"
	// "strconv"

	// "io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Roll string
}

func check(w http.ResponseWriter, r *http.Request) {

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatalln(err.Error())

	}

	roll, err := strconv.Atoi(p.Roll)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	coin := 0.00
	db, err := sql.Open("sqlite3", "./a.db")
	if err != nil {
		log.Fatalln(err.Error())
		return

	}
	defer db.Close()

	sqlStatement := `SELECT coins FROM User WHERE id=$1;`
	row, err := (db.Query(sqlStatement, roll))
	if err != nil {
		log.Fatalln(err.Error())
		return

	}

	responses := 0
	for row.Next() {
		_ = row.Scan(&coin)
		responses++
	}

	if responses == 0 {
		fmt.Fprintf(w, "USER DOESNT EXIST:: goto /signup to create users \n")
		return
	}

	fmt.Fprintf(w, "Balance : %d\n", int(coin))

}
