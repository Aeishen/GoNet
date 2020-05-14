package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func server(ws *websocket.Conn) {
	fmt.Printf("new connection\n")
	_, _ = io.WriteString(ws, "i am here")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
	fmt.Printf(" => closing connection\n")
	_ = ws.Close()
}

func main() {
	http.Handle("/websocket", websocket.Handler(server))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
