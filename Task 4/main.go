package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3" 

)



func main() {
 
	dbase, _ := sql.Open("sqlite3", "./a.db")
// ?	tx, _ := sql.Open("sqlite3", "./transactions.db")

	stat, _ := dbase.Prepare("CREATE TABLE IF NOT EXISTS  User (id INTEGER PRIMARY KEY, coins FLOAT, pass TEXT, adm INTEGER)")
	stat.Exec()
	stat2, err := dbase.Prepare("CREATE TABLE IF NOT EXISTS  Tx (id INTEGER PRIMARY KEY,Time TEXT , coins INTEGER, txfrom INTEGER, txto INTEGER)")
	
	if(err!=nil){log.Fatal(err.Error())}
	 stat2.Exec()


	http.HandleFunc("/create/", create)
	http.HandleFunc("/transfer/", transfer)
	http.HandleFunc("/check/", check)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/signup/", signup)


	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}
