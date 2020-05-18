//目前的jsonrpc库是基于tcp协议实现的，暂时不支持使用http进行数据传输
package main

import (
	"./rpcObjects"
	"io"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	calc := new(rpcObjects.Args) // 服务器创建一个用于计算的对象

	_ = rpc.Register(calc)        // 注册rpc服务

	listener, e := net.Listen("tcp", "localhost:1234")  // 开启监听
	if e != nil {
		log.Fatalln("Starting RPC-server -listen error:", e)
	}
	_, _ = io.WriteString(os.Stdout, "start connection\n")

	for {
		conn, err := listener.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			_, _ = io.WriteString(os.Stdout, "new client in coming\n")
			jsonrpc.ServeConn(conn)  // 跟originalRpc一样, 最终都是调用RPC Server的 ServeCodec 方法
		}(conn)
	}
}
