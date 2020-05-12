// 简单的多线程多核 TCP server.

package main

import (
	"GoNet/goWeb"
	"errors"
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25
var _ map[string]bool  // todo cliList

func main() {

	// 获取命令行输入
	flag.Parse()
	if flag.NArg() != 2{
		goWeb.ErrorHandle(errors.New("error flag args"),"")
	}


	// 格式化成字符串
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))

	//
	listener := initServer(hostAndPort)

	for{
		conn, err := listener.Accept()
		goWeb.ErrorHandle(err,"listener Accept err...")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) (listener net.Listener){
	var err error
	var tcpAddr *net.TCPAddr

	// ResolveTCPAddr返回TCP端点的地址
	tcpAddr, err = net.ResolveTCPAddr("tcp", hostAndPort)
	goWeb.ErrorHandle(err,"initServer ResolveTCPAddr err...")

	listener, err = net.ListenTCP("tcp", tcpAddr)
	goWeb.ErrorHandle(err,"initServer ListenTCP err...")

	fmt.Println("Listening to: ", listener.Addr().String())
	return
}

func connectionHandler(conn net.Conn){

	//获取客户端的地址
	connFrom := conn.RemoteAddr().String()

	println("Connection from: ", connFrom)

	sayToCli(conn)

	for{
		inBuf := make([]byte, maxRead+1)
		length, err := conn.Read(inBuf[0:maxRead])
		inBuf[maxRead] = 0 // 作为输入结束符, 防止溢出


		switch err {
		case nil:
			handleMsg(length, inBuf)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT //出现错误就关闭连接
		}
	}

DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	goWeb.ErrorHandle(err,"Close: ")
}

func sayToCli(conn net.Conn){
	outBuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := conn.Write(outBuf)
	goWeb.ErrorHandle(err,"Write: wrote "+string(wrote)+" bytes.")
}


func handleMsg(length int, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">\n")
	}
}
