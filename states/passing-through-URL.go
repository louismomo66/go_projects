package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000",nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	v := r.FormValue("q") // Returns the first value for the named component of the query
	io.WriteString(w, "Do my search: "+v) //Retrieve the value from a URL
}

//Example
//http://localhost:9000/?q=John