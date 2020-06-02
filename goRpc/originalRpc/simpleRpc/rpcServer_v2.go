package main

import (
	"./rpcObjects"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	_ = rpc.RegisterName("Args", new(rpcObjects.Args)) //通过名称注册对象, 底层调用跟rpc.Register一样,只是有无名字区别

	listener, e := net.Listen("tcp", ":1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	go rpc.Accept(listener) //在侦听器上接受连接，并为每个传入连接提供请求

	// 下面代码就是rpc.Accept的内部实现
	//for {
	//	conn, e := listener.Accept() // 接收客户端连接请求
	//	if e != nil {
	//		log.Fatal("Starting RPC-server -listen Accept:", e)
	//		continue
	//	}
	//	go rpc.ServeConn(conn)
	//}
	time.Sleep(1000e9)
}