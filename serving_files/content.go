package main

import (
	"io"
	"net/http"
	"os"
)

func d(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request){
	f,err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found ",404)
		return
	}
	defer f.Close()

	fi,err := f.Stat() //give pointer to a file and then file info eg name, modification time,size,mode
	if err != nil{
		http.Error(w, "file not found",404)
		return
	}
	http.ServeContent(w,r, f.Name(), fi.ModTime(),f) /*
	func ServeContent(w ResponseWrite, r*Request, name string, modtime time.Time, content io.ReadSeeker)
	*/
}

func main(){
	http.HandleFunc("/",d)
	http.HandleFunc("/toby.jpg",dogPic)
	http.ListenAndServe(":9000",nil)
}