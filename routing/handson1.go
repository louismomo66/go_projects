package main

import (
	"io"
	"net/http"
)

func root(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "This is the root")
}
func d(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "This is a dog")
}
func me(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "This is me")
}

func main(){
	http.HandleFunc("/",root)
	http.HandleFunc("/dog/",d)
	http.HandleFunc("/me/",me)

	http.ListenAndServe(":9000",nil)
}