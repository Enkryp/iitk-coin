package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3" 

)



func main() {
 
	dbase, _ := sql.Open("sqlite3", "./a.db")
	stat, _ := dbase.Prepare("CREATE TABLE IF NOT EXISTS  User (id INTEGER PRIMARY KEY, coins INTEGER)")
	stat.Exec()


	http.HandleFunc("/create/", create)
	http.HandleFunc("/transfer/", transfer)
	http.HandleFunc("/check/", check)


	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}
