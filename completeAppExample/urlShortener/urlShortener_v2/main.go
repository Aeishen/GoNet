package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var myStore *URLStore

func Add(w http.ResponseWriter, r *http.Request)  {
	url := r.FormValue("url")
	if url == ""{
		w.Header().Set("content-type", "text/html")
		// fmt.Fprint也可以,但io.WriteString速度更快, 因为前者接受的是切片的接口，后者只是一个字符串参数。
		_, _ = io.WriteString(w, AddForm)
		//_, _ = fmt.Fprint(w, AddForm)
		return
	}
	key := myStore.Put(url)
	_, _ = io.WriteString(w, "http://localhost:8080/" + key)
	//_, _ = fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

// 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := myStore.Get(key)

	if url == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}


func main() {
	myStore = NewURLStore("store.gob")
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatalf("URl Start Http Listen : %v", err)
	}
}