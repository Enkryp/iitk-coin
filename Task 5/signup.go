package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type signup_req struct {
	Roll string
	Pass string
	OTP  string
}

func signup(w http.ResponseWriter, r *http.Request) {

	var p signup_req

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatalln(err.Error())

	}

	roll, err := strconv.Atoi(p.Roll)
	if err != nil {
		log.Fatalln(err.Error())
		return

	}

	pass := (p.Pass)
	dbase, _ := sql.Open("sqlite3", "./a.db")

	defer dbase.Close()

	// 	//// otp

	if otp(dbase, roll, p, w) == 1 {
		dbase.Close()

		dbase, _ = sql.Open("sqlite3", "./a.db")
		enter(dbase, roll, pass, w)
	}

}

func otp(dbase *sql.DB, roll int, p signup_req, w http.ResponseWriter) int {

	ex := `SELECT pass FROM User WHERE id=$1;`
	rows, err := dbase.Query(ex, roll)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if rows.Next() {
		fmt.Fprint(w, "USER ALREADY PRESENT ")
		for rows.Next() {
		}
		return 0
	}

	layout := time.Now().String()

	if p.OTP == "NULL" {

		ex = `SELECT id FROM Otp WHERE id=$1;`
		rows, err = dbase.Query(ex, roll)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if rows.Next() {
			fmt.Fprint(w, "Otp Already generated; Please enter now\n")
			for rows.Next() {
			}
			return 0
		}

		nBig, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			log.Fatalln(err.Error())
		}
		n := nBig.Int64()
		ex := `INSERT INTO Otp (id,fail,pass,time) VALUES (?,?,?,?) `
		sqlcode, err := dbase.Prepare(ex)

		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = sqlcode.Exec(roll, 0, n, time.Now())
		if err != nil {
			log.Fatalln(err.Error())
		}
		mail(strconv.Itoa(roll)+"@iitk.ac.in", strconv.Itoa(int(n)))
		fmt.Fprintf(w, "Now Enter OTP")
		return 0

	}

	ex = `SELECT time,fail,pass FROM Otp WHERE id=$1;`
	rows, err = dbase.Query(ex, roll)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if !rows.Next() {
		fmt.Fprint(w, "Please gen OTP first")
		return 0
	}
	for rows.Next() {
	}

	ex = `SELECT time,fail,pass FROM Otp WHERE id=$1;`
	rows, err = dbase.Query(ex, roll)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var time2 string
	var fail int
	var otp string

	for rows.Next() {
		err = rows.Scan(&time2, &fail, &otp)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}

	time_otp, _ := time.Parse(layout, time2)
	time_exp := time_otp.Add(5 * time.Minute)

	if time_exp.After(time.Now()) || fail == 2 {

		fmt.Fprint(w, "Expired/ 3rd Try !\n Generate again \n")
		ex = `DELETE FROM Otp WHERE id=$1;`
		sqlcode, err := dbase.Prepare(ex)

		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = sqlcode.Exec(roll)
		if err != nil {
			log.Fatalln(err.Error())
		}

		return 0

	}

	if p.OTP != otp {

		fmt.Fprint(w, "Wrong Otp!")

		ex = `UPDATE Otp SET fail= fail+1 WHERE id=$1;`
		sqlcode, err := dbase.Prepare(ex)

		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = sqlcode.Exec(roll)
		if err != nil {
			log.Fatalln(err.Error())
		}

		return 0

	}

	// fmt.Println("HERE2")
	dbase.Close()

	dbase, _ = sql.Open("sqlite3", "./a.db")
	ex = `DELETE FROM Otp WHERE id=$1;`
	sqlcode2, err := dbase.Prepare(ex)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// fmt.Println("HERE5")

	_, err = sqlcode2.Exec(roll)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// fmt.Println("HERE1")

	return 1
}

func enter(db *sql.DB, rollno int, pass string, w http.ResponseWriter) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

	hash := string(hashed)
	ex := `SELECT pass FROM User WHERE id=$1;`
	rows, _ := db.Query(ex, rollno)

	if rows.Next() {
		fmt.Fprint(w, "USER ALREADY PRESENT ")
		for rows.Next() {
		}
		return
	}

	new := `INSERT INTO User(id, pass,coins,adm) VALUES (?, ?, ?,?)`

	sqlcode, err := db.Prepare(new)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(rollno, hash, 0, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintf(w, "Done !\n")
}
