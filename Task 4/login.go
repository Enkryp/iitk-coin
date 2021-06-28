package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	// "time"

	// "fmt"
	"strconv"

	// "log"

	// "reflect"

	// "fmt"
	// "html/template"
	// "log"
	"net/http"
	// "strconv"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// "io/ioutil"
)

// func logout(w http.ResponseWriter, r *http.Request) {
// // return
// // var cookie http.Cookie
// // cookie.Name="JWT"
// // cookie.Value="nil"
// // cookie.Path="/"
// // http.SetCookie(w,&cookie)
// // // fmt.Println("NULLED")
// // fmt.Fprint(w,"LOGGED OUT")

// }

type login_req struct {
    Roll string
	Pass string
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("signup access")

		
	var p login_req

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

	auth(roll, pass,w,r)

	}




func auth(user int , pass string, w http.ResponseWriter, r *http.Request) {

	roll := user
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
	if(hi==0){fmt.Fprintf(w, "USER DOESNT EXIST"); return}


	bpass := []byte(pass2)
	

	// fmt.Println([]byte(pass))
	err := bcrypt.CompareHashAndPassword(bpass, []byte(pass))
	// a,_:=bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	// fmt.Println(pass2, string(a))
    if err != nil {
		
		fmt.Fprintf(w,"Wrong Cred!")
		return
    }
    
    

alpha:= strconv.Itoa(user)
head:= base64.URLEncoding.EncodeToString([]byte(`{"alg":"HS256", "typ":"JWT"}`))
payload:= base64.URLEncoding.EncodeToString([]byte(`"User": ` + `"`+alpha+ `"`))
body := head +`.` +payload



key := []byte("yoyoyoyoyoyo")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(body))
signature:=base64.StdEncoding.EncodeToString(h.Sum(nil))

JWT:= body+`.`+signature


// var cookie http.Cookie
// cookie.Name="JWT"
// cookie.Value=JWT
// cookie.Path="/"
// http.SetCookie(w,&cookie)

// fmt.Println("SET", cookie.Value)



fmt.Fprintf(w,"Your JWT : " + JWT+ "\n")

	

}
