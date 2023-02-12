package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main(){
	http.HandleFunc("/",foo)
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":9000",nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("sessiom-id")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w,cookie)
	}
	fmt.Println(cookie)
}
