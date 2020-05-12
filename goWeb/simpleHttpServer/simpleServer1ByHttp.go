package main

import (
	"GoNet/goWeb"
	"fmt"
	"io"
	"net/http"
)

// 使用处理器
type HttpHandler struct {}

func (h *HttpHandler)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("Inside httpHandler")
	//_, _ = fmt.Fprintf(w, "HttpHandler : Hello,"+req.URL.Path[1:]) // 写法1
	_, _ = io.WriteString(w, "HttpHandler: Hello,"+req.URL.Path[1:]) // 写法2
}

// 使用处理函数 该函数必须实现与http.Handler接口下的ServeHTTP方法同样的签名
func ServeHTTPFunc(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("Inside ServeHTTPFunc")
	//_, _ = fmt.Fprintf(w, "ServeHTTPFunc : Hello,"+req.URL.Path[1:])
	_, _ = io.WriteString(w, "ServeHTTPFunc: Hello,"+req.URL.Path[1:])
}

func main() {
	httpHandler := new(HttpHandler)

	// 第一个参数是请求的路径，第二个参数是处理这个路径请求的函数的引用，有如下多种方法注册处理
	http.Handle("/",httpHandler) // 使用处理器
	//http.HandleFunc("/",ServeHTTPFunc) // http.HandleFunc 注册了一个处理函数 (这里是 ServeHTTPFunc) 来处理对应 / 的请求。
	//http.Handle("/",http.HandlerFunc(ServeHTTPFunc)) // HandlerFunc 实现了 Handler 的ServeHTTP接口, 所以HandlerFunc 也算是一个 Handler

	// 与net.Listen("tcp", "localhost:8080")类似, 前面已经注册了处理器,所以此处handler为nil
	err := http.ListenAndServe("localhost:8080", nil)

	// 以上可以替换为下面两种写法, HandlerFunc 实现了 Handler 的ServeHTTP接口, 所以HandlerFunc 也算是一个 Handler
	//err := http.ListenAndServe("localhost:8080", httpHandler)
	//err := http.ListenAndServe("localhost:8080", http.HandlerFunc(ServeHTTPFunc))

	goWeb.ErrorHandle(err, "http.ListenAndServe err")
}
