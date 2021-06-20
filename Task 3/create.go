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

type create_req struct {
    Roll string
	Coins string
}

func create (w http.ResponseWriter, r *http.Request){


	
	var p create_req

    err := json.NewDecoder(r.Body).Decode(&p)
	if(err!=nil){
		log.Fatalln(err.Error())

	}



	roll,err:=strconv.Atoi(p.Roll)
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	coin,err :=strconv.Atoi(p.Coins)
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	db, err:= sql.Open("sqlite3", "./a.db")
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}
	defer db.Close()

	sqlStatement := `SELECT coins FROM User WHERE id=$1;`
	row,_ := (db.Query(sqlStatement, roll))
	
	responses:=0
	for row.Next(){		responses++;	}


	if(responses==0){
		


	new := `INSERT INTO User(id, coins) VALUES (?, ?)`

	sqlcode, err := db.Prepare(new)


	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	_, err = sqlcode.Exec(roll, coin)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
		
	}	else {


		new:="UPDATE User SET  coins = coins + $1 WHERE id = $2;"
		sqlcode, err := db.Prepare(new)

		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		_, err = sqlcode.Exec(coin,roll)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

	}

	fmt.Fprintf(w,"Done!\n")

}