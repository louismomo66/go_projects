package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
http.HandleFunc("/",set)
http.HandleFunc("/read",read)
http.HandleFunc("/aboundance",aboundace)
http.Handle("/favicon.ico",http.NotFoundHandler())
http.ListenAndServe(":9000",nil)

}

func set(w http.ResponseWriter, r *http.Request){
       http.SetCookie(w, &http.Cookie{
				Name:"mycookie",
				Value: "some-value",
			 })
			 fmt.Fprintln(w, "COOKIE WRITEN - CHECK YOUR BROWSER")
			 fmt.Fprintln(w, "in chrom go to; dev tools / application / cookies")

}

func read(w http.ResponseWriter, r *http.Request){
	c1,err := r.Cookie("mycookie")
	if err != nil{
		log.Println(err)
	}else{
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2,err := r.Cookie("specific")
	if err != nil{
		log.Println(err)
	}else{
		fmt.Fprintln(w, "YOUR COOKIE #1:", c2)
	}

	c3,err := r.Cookie("general")
	if err != nil{
		log.Println(err)
	}else{
		fmt.Fprintln(w, "YOUR COOKIE #1:", c3)
	}
}

func aboundace(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "specific",
		Value: "some other value about specific",

	})

	http.SetCookie(w, &http.Cookie{
		Name: "general",
		Value: "some other value about general",

	})

	fmt.Fprintln(w, "COOKIEs WRITEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrom go to; dev tools / application / cookies")
}