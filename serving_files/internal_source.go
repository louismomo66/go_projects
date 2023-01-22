package main

import (
	"io"
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", d)
http.HandleFunc("/toby.jpg",dogPic)
http.ListenAndServe(":9000",nil)
}

func d(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src="/toby.jpg">
	`) //write to the resposewriter
}

func dogPic(w http.ResponseWriter, r *http.Request){
	f,err := os.Open("toby.jpg") //give pointer to a file
	if err != nil{
		http.Error(w, "file not found",404) //handle any error
		return
	}
	defer f.Close()
	io.Copy(w,f) // copy from the file to the responseWriter
}
