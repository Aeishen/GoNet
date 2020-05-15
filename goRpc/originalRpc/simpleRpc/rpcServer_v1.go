//采用http协议作为rpc载体

package main

import (
	"./rpcObjects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpcObjects.Args) // 服务器创建一个用于计算的对象

	_ = rpc.Register(calc)        // 注册rpc服务

	rpc.HandleHTTP()              // HandleHTTP在默认rpcPath上注册用于RPC消息的HTTP处理程序，并在默认debugPath上注册调试处理程序

	listener, e := net.Listen("tcp", "localhost:1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	// 对每一个进入到 listener 的请求，都是由协程去启动一个 http.Serve(listener, nil) ， 为每一个
	// 传入的 HTTP 连接创建一个新的服务线程。同时我们必须保证在一个特定的时间内服务器是唤醒状态
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}