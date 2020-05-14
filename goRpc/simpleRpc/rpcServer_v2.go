package main

import (
	"./rpcObjects"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	_ = rpc.RegisterName("AArgs", new(rpcObjects.Args))  //通过名称注册对象

	listener, e := net.Listen("tcp", ":1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	conn, e := listener.Accept()
	if e != nil {
		log.Fatal("Starting RPC-server -listen Accept:", e)
	}

	rpc.ServeConn(conn)
	time.Sleep(1000e9)
}