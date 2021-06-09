package main

import (
	"database/sql"
	// "fmt"
	"fmt"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// "io/ioutil"
)





func signup (w http.ResponseWriter, r *http.Request){
	// fmt.Println("signupdone called")
	t, _ := template.ParseFiles("signup.html")
	t.Execute(w, "signup")

	user:= r.FormValue("user")
	pass:= r.FormValue("pass")
	if(user=="" || pass==""){return}
	dbase, _ := sql.Open("sqlite3", "./a.db")
	stat, _ := dbase.Prepare("CREATE TABLE IF NOT EXISTS  User (id INTEGER PRIMARY KEY, pass TEXT)")
	stat.Exec()

	defer dbase.Close()
	a,_:= strconv.Atoi(user)
	enter(dbase, a, pass,w)
	
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

	new := `INSERT INTO User(id, pass) VALUES (?, ?)`

	sqlcode, err := db.Prepare(new)


	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(rollno, hash)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintf(w,"DONE")
}

