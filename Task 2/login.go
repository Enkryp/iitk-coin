package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	// "time"

	// "fmt"
	"strconv"

	// "log"

	// "reflect"

	// "fmt"
	"html/template"
	// "log"
	"net/http"
	// "strconv"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// "io/ioutil"
)

func logout(w http.ResponseWriter, r *http.Request) {
// return
var cookie http.Cookie
cookie.Name="JWT"
cookie.Value="nil"
cookie.Path="/"
http.SetCookie(w,&cookie)
// fmt.Println("NULLED")
fmt.Fprint(w,"LOGGED OUT")

}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("signup access")

	if(r.FormValue("logout")=="true"){logout(w,r); return}
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	// fmt.Println(user, pass)
	if user != "" && pass != "" {

		auth(user, pass,w,r)

	}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, "login")

}

func auth(user, pass string, w http.ResponseWriter, r *http.Request) {

	roll, _ := strconv.Atoi(user)
	db, _ := sql.Open("sqlite3", "./a.db")
	defer db.Close()
	var pass2 string
	
	sqlStatement := `SELECT pass FROM User WHERE id=$1;`
	row,_ := (db.Query(sqlStatement, roll))
	hi:=0
	for row.Next(){
		_= row.Scan(&pass2)
		hi++
		
	}
	if(hi==0){fmt.Fprintf(w, "<html>USER DOESNT EXIST</html>"); return}


	bpass := []byte(pass2)
	

	// fmt.Println([]byte(pass))
	err := bcrypt.CompareHashAndPassword(bpass, []byte(pass))
	// a,_:=bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	// fmt.Println(pass2, string(a))
    if err != nil {
		t, _ := template.ParseFiles("wrong.html")
	t.Execute(w, "sryy")
        logout(w,r)
		// fmt.Println(err.Error())
		return
    }
    
    


head:= base64.URLEncoding.EncodeToString([]byte(`{"alg":"HS256", "typ":"JWT"}`))
payload:= base64.URLEncoding.EncodeToString([]byte(`"user":` + string(user)))
body := head +`.` +payload



key := []byte("yoyoyoyoyoyo")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(body))
signature:=base64.StdEncoding.EncodeToString(h.Sum(nil))

JWT:= body+`.`+signature


var cookie http.Cookie
cookie.Name="JWT"
cookie.Value=JWT
cookie.Path="/"
http.SetCookie(w,&cookie)

// fmt.Println("SET", cookie.Value)





	
	t, _ := template.ParseFiles("ok.html")
	t.Execute(w, "Yipee")

}
