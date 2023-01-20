package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func root(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "This is the root")
}
func d(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "This is a dog")
}
func me(res http.ResponseWriter, r *http.Request){
	io.WriteString(res, "Kwezi")
}

func mycloud(res http.ResponseWriter, r *http.Request){
tpl,err := template.ParseFiles("something.gohtml")
	if err != nil{
  log.Fatalln("error trying to parse template",err)
	}
	
err = tpl.ExecuteTemplate(res,"something.gohtml","Kwezi")
if err != nil{
	log.Fatalln("erro tring to execute template")
}
}






func main(){
	http.HandleFunc("/",root)
	http.HandleFunc("/dog/",d)
	http.HandleFunc("/me/",mycloud)

	http.ListenAndServe(":9000",nil)
}
