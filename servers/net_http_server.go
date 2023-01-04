package main

import (
	"fmt"
	"net/http"
	 
)

type handler int

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Any code you want in this func")
}

func main(){
	var d handler
	http.ListenAndServe(":9000",d)
}