package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"fmt"
	// "fmt"
	// "html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// "io/ioutil"
)



type signup_req struct {
    Roll string
	Pass string
}

func signup (w http.ResponseWriter, r *http.Request){
	// fmt.Println("signupdone called")
	// fmt.Println("In")


		
	var p signup_req

    err := json.NewDecoder(r.Body).Decode(&p)
	if(err!=nil){
		log.Fatalln(err.Error())

	}


	roll,err:=strconv.Atoi(p.Roll)
	if(err!=nil){
		log.Fatalln(err.Error())
		return

	}

	pass:=(p.Pass)
	dbase, _ := sql.Open("sqlite3", "./a.db")
	// stat, _ := dbase.Prepare("CREATE TABLE IF NOT EXISTS  User (id INTEGER PRIMARY KEY, pass TEXT)")
	// stat.Exec()

	defer dbase.Close()
	enter(dbase, roll, pass,w)
	
}





func enter(db *sql.DB, rollno int, pass string, w http.ResponseWriter) {

	// fmt.Println("USER %d PASS %s",rollno, pass)
	// fmt.Println([]byte(pass))
	hashed,_:= bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost) 

	hash:=string(hashed)
	// print(hash)
	// err:= bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	// if(err!=nil){fmt.Println(err.Error())}
	ex := `SELECT pass FROM User WHERE id=$1;`
	rows,_:= db.Query(ex,rollno)

	if(rows.Next()){fmt.Fprint(w,"USER ALREADY PRESENT "); return}

	new := `INSERT INTO User(id, pass,coins,adm) VALUES (?, ?, ?,?)`

	sqlcode, err := db.Prepare(new)


	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(rollno, hash,0,0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintf(w,"Done !\n")
}

