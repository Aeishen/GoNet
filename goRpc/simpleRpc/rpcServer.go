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
	calc := new(rpcObjects.Args)  // 服务器创建一个用于计算的对象

	_ = rpc.Register(calc)        // 注册进rpc
	//_ = rpc.RegisterName("Calculator", calc)  //也可以通过名称注册对象

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", "localhost:1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	// 对每一个进入到 listener 的请求，都是由协程去启动一个 http.Serve(listener, nil) ， 为每一个
	// 传入的 HTTP 连接创建一个新的服务线程。同时我们必须保证在一个特定的时间内服务器是唤醒状态
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}