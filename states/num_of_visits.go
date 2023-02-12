package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/",visits)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000",nil)
}

func visits(w http.ResponseWriter, r *http.Request){
	cookie,err := r.Cookie("my-cookie")

	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name: "my-cookie",
			Value: "0",
		}
	}
 count, err := strconv.Atoi(cookie.Value)
 if err != nil{
	log.Fatalln(err)
 }
 count++
 cookie.Value = strconv.Itoa(count)
 http.SetCookie(w, cookie)
 io.WriteString(w, cookie.Value)

}