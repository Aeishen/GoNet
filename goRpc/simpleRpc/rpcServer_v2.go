package main

import (
	"./rpcObjects"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	_ = rpc.RegisterName("Args", new(rpcObjects.Args))  //通过名称注册对象

	listener, e := net.Listen("tcp", ":1234")  // 开启监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	go func() {
		for {
			conn, e := listener.Accept()
			if e != nil {
				log.Fatal("Starting RPC-server -listen Accept:", e)
			}
			go rpc.ServeConn(conn)
		}
	}()
	time.Sleep(1000e9)
}