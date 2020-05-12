package main

import (
	"GoNet/goWeb"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	cli, err := net.Dial("tcp", "localhost:8080")
	goWeb.ErrorHandle(err," cli Dial err")

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("首先, 输入你的名字?")

	clientName, err := inputReader.ReadString('\n')
	goWeb.ErrorHandle(err," cli ReadString err")

	trimmedClient := strings.Trim(clientName, "\r\n") // "\r\n" on Windows, "\n" on Linux
	_, err = cli.Write([]byte(trimmedClient + " coming" ))
	goWeb.ErrorHandle(err," cli Write err")


	for{
		fmt.Println("你想发送什么给服务器? (输入Q可以退出)")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")

		cli.Write([]byte(trimmedClient + " says: " + trimmedInput))

		if trimmedInput == "Q" {
			return
		}
	}
}
