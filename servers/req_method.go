package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type handler int

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
err := r.ParseForm()
if err != nil{
log.Fatalln(err)
}

data := struct{
	Method string
	Submissions url.Values
}{
	r.Method,
	r.Form,
}
tpl.ExecuteTemplate(w, "index4.gohtml", data)

}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index4.gohtml"))
}
func main(){
	var d  handler
	http.ListenAndServe(":9000",d)
}
