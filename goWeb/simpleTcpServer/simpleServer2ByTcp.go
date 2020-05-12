package main

import (
	"GoNet/goWeb"
	"fmt"
	"net"
	"os"
	"strings"
)

// 服务器上存储的连接过的客户端的状态 1代表活跃, 0代表下线
var mapUsers map[string]int

func main() {
	var listener net.Listener
	var err error
	var conn net.Conn
	mapUsers = make(map[string]int)
	fmt.Println("Starting the server ...")

	listener, err = net.Listen("tcp", "localhost:8080")
	goWeb.ErrorHandle(err, "net.Listen err...")

	for {
		conn, err = listener.Accept()
		goWeb.ErrorHandle(err, "listener.Accept err...")
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var err error
	var clName string

	for {
		// 获得客户端输入
		input := string(buf)

		// 获取客户端名字:
		if clName == ""{
			clName = input
		}

		// 设置客户端为活跃状态
		mapUsers[clName] = 1

		buf = make([]byte, 512)
		_, err = conn.Read(buf)
		goWeb.ErrorHandle(err, "conn.Read err...")

		fmt.Printf("Received data: --%v\n--", string(buf))

		// 关闭服务器
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}

		// 打印所有客户端活跃状态
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}

		// 打印所有客户端活跃状态
		if strings.Contains(input, ": Q") {
			mapUsers[clName] = 0
			fmt.Printf("%s levaing...\n", clName)
		}
	}
}

func DisplayList() {
	fmt.Println("||---------------------------------------------||")
	fmt.Println("||This is the client list: 1=active, 0=inactive||")
	for key, value := range mapUsers {
		fmt.Printf("||--------------User %s is %d\n---------------||", key, value)
	}
	fmt.Println("||---------------------------------------------||")
}