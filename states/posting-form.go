package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",poster)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

func poster(w http.ResponseWriter, r *http.Request){
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
`<form method="GET">
<input type="text" name="q"
<input type="submit">
</form>
<br>` +v)
}

//post: The value goes through the body
//get: The value goes through the URL