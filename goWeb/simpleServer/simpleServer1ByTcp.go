package main

import (
	"GoNet/goWeb"
	"fmt"
	"net"
)

func main() {

	// 创建 listener
	ser, err := net.Listen("tcp" ,"localhost:8080")
	goWeb.ErrorHandle(err," ser Listen err")

	// 监听并接受来自客户端的连接
	for {
		conn, err := ser.Accept()
		goWeb.ErrorHandle(err," ser Accept err")

		// 处理来自客户端的请求
		go handleConn1(conn)
	}
}

func handleConn1(conn net.Conn)  {
	for{
		buf := make([]byte,512)
		_, err := conn.Read(buf)
		goWeb.ErrorHandle(err," ser Read err")
		fmt.Println("client say: ", string(buf))
	}
}
