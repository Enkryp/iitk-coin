package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	// "database/sql"
	// "os"
	// _ "github.com/mattn/go-sqlite3"
)
func def (w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "MAIN")
}

func main() {
 

	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/secret/", secret)
	http.HandleFunc("/", def)


	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}

