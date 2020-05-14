// 网页服务器1

package main

import (
	"GoNet/goWeb"
	"io"
	"net/http"
)

// 一个文本框和一个提交按钮
const form2 = `<html><body><form action="#" method="post" name="bar">
		      <input type="text" name="in"/>
			  <input type="submit" value="Submit"/>
			  </form></html></body>`

func SimpleServer(w http.ResponseWriter, req *http.Request)  {
	_, err := io.WriteString(w, "<h1>hello, world</h1>")
	goWeb.ErrorHandle(err,"SimpleServer :")
}

func FromServer(w http.ResponseWriter, req *http.Request){

	// 在写入返回之前将 header 的 content-type 设置为 text/html, content-type 会让浏览器认为它可以使用函数 http.DetectContentType([]byte(form)) 来处理收到的数据
	w.Header().Set("Content-Type","text/html")

	switch req.Method {
	case "GET":
		_, err := io.WriteString(w, form2)
		goWeb.ErrorHandle(err,"FromServer GET:")
	case "POST":

		// 使用 request.FormValue("in") 通过文本框的 name 属性 in 来获取内容,并写回浏览器页面
		_, err := io.WriteString(w, req.FormValue("in"))
		goWeb.ErrorHandle(err,"FromServer POST:")
	}
}

func main() {

	// 使用LogPanics处理panic
	http.HandleFunc("/simple", LogPanics(SimpleServer))
	http.HandleFunc("/from", LogPanics(FromServer))

	err := http.ListenAndServe("localhost:8080", nil)
	goWeb.ErrorHandle(err,"ListenAndServe :")
}