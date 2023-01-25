package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct{
	FirstName string
	LastName string
	Subscribed bool
}

func main(){
	http.HandleFunc("/",poster)
	http.Handle("/favicon.ico",http.NotFoundHandler())
  http.ListenAndServe(":9000",nil)
}

func poster(w http.ResponseWriter, r *http.Request ){
f := r.FormValue("first")
l := r.FormValue("second")
s := r.FormValue("subscribes") == "on"

err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})

if err != nil{
	http.Error(w, err.Error(), 500)
	log.Fatalln(err)
}




}