package main

import (
	"GoNet/goRpc/protoRpc/simpleRpc/rpcObjects"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpcObjects.Args) // 服务器创建一个用于计算的对象

	_ = rpc.Register(calc)        // 注册rpc服务

	listener, e := net.Listen("tcp", ":1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	go rpc.Accept(listener) //在侦听器上接受连接，并为每个传入连接提供请求
	time.Sleep(1000e9)
}
