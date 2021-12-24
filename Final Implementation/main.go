package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	dbase, _ := sql.Open("sqlite3", "./a.db")
	// ?	tx, _ := sql.Open("sqlite3", "./transactions.db")

	stat, _ := dbase.Prepare("CREATE TABLE IF NOT EXISTS  User (id INTEGER PRIMARY KEY, coins FLOAT, pass TEXT, adm INTEGER)")
	stat.Exec()
	stat2, err := dbase.Prepare("CREATE TABLE IF NOT EXISTS  Tx (id INTEGER PRIMARY KEY,Time TEXT , coins INTEGER, txfrom TEXT, txto INTEGER)")

	if err != nil {
		log.Fatal(err.Error())
	}
	stat2.Exec()
	stat3, err := dbase.Prepare("CREATE TABLE IF NOT EXISTS  Redeem (id INTEGER PRIMARY KEY,Item TEXT , coins INTEGER, Recipient INTEGER)")

	if err != nil {
		log.Fatal(err.Error())
	}
	stat3.Exec()

	stat4, err := dbase.Prepare("CREATE TABLE IF NOT EXISTS  Pending (id INTEGER PRIMARY KEY,awardId INTEGER, coins INTEGER, Recipient INTEGER)")

	if err != nil {
		log.Fatal(err.Error())
	}
	stat4.Exec()

	stat5, err := dbase.Prepare("CREATE TABLE IF NOT EXISTS  Otp (id INTEGER PRIMARY KEY, pass TEXT, time TEXT, fail INTEGER)")

	if err != nil {
		log.Fatal(err.Error())
	}
	stat5.Exec()

	http.HandleFunc("/create/", create)
	http.HandleFunc("/transfer/", transfer)
	http.HandleFunc("/check/", check)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/redeem/", redeem)
	http.HandleFunc("/add/", add)
	http.HandleFunc("/approve/", approve)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
