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

	output := make([]byte, 25)
	_, err = cli.Read(output)
	goWeb.ErrorHandle(err," cli dealRespond err")
	fmt.Println("server say: ",string(output))

	for{
		fmt.Println("你想发送什么给服务器? (输入Q可以退出)")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")

		_, _ = cli.Write([]byte(trimmedClient + " says: " + trimmedInput))

		if trimmedInput == "Q" {
			return
		}


	}
}
