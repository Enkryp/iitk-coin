package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func secret(w http.ResponseWriter, r *http.Request) {


y,err:= r.Cookie("JWT")
if err!=nil{

	fmt.Fprint(w,"SORRY")
	return
}
a:=y.Value
// fmt.Println(a)
if(len(a)<4){fmt.Fprint(w,"SORRY"); return}

x:= strings.Split(a,".")
// sig:= strings.Split(x[2],";")[0]

body:= x[0]+"." +x[1]
sig:= x[2]


key := []byte("yoyoyoyoyoyo")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(body))
signature:=base64.StdEncoding.EncodeToString(h.Sum(nil))

if(sig==signature){
	fmt.Fprint(w,"Congo, your sec num is 128")
	return
}
fmt.Fprintf(w,"Sorry auth first")



}