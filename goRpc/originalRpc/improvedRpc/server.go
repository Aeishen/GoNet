package main

import (
	"GoNet/goRpc/originalRpc/improvedRpc/api"
	"GoNet/goRpc/originalRpc/improvedRpc/rpcObjects"
	"log"
	"net"
	"net/rpc"
)

// 该服务也满足ArgsServiceInterface接口
type ArgsServer struct {}

func (a *ArgsServer) Multiply(req *rpcObjects.Args, resp *int) error {
	*resp = req.M * req.N
	return nil
}

func (a *ArgsServer) Add(req *rpcObjects.Args, resp *int) error {
	*resp = req.M + req.N
	return nil
}

func main() {
	api.RegisterArgsService(new(ArgsServer))
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Starting RPC-server -listen error:", err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal("listener.Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
