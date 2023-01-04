package main

import (
	"html/template"
	"log"
	"net/http"
)

type handler int

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w,"index3.gohtml",r.Form)
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index3.gohtml"))
}

func main(){
	var d handler
	http.ListenAndServe(":9000",d)
}