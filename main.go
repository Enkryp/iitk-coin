package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {

	os.Remove("a.db")

	file, err := os.Create("a.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	file.Close()

	dbase, _ := sql.Open("sqlite3", "./a.db")
	defer dbase.Close()

	table(dbase)

	for i := 1000; i < 1050; i++ {

		enter(dbase, i, "abc"+strconv.Itoa(i))

	}

	disp(dbase)

	//comment to not display
}

func table(db *sql.DB) {

	tab := `CREATE TABLE User (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,	
		"rollno" integer,	
		"name" TEXT
				
	  );`

	sqlcode, err := db.Prepare(tab)
	log.Println("dbg")

	if err != nil {
		log.Fatal(err.Error())
	}
	sqlcode.Exec()
}

func enter(db *sql.DB, rollno int, name string) {

	new := `INSERT INTO User(rollno, name) VALUES (?, ?)`

	sqlcode, err := db.Prepare(new)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = sqlcode.Exec(rollno, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func disp(db *sql.DB) {

	row, err := db.Query("SELECT * FROM User ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var rollno int
		var name string
		row.Scan(&id, &rollno, &name)
		log.Println("Roll_no: ", rollno, " Name:", name)
	}
}
