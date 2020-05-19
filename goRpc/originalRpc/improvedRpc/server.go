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
	rpc.Accept(listener)
	//time.Sleep(1000e9)
}
