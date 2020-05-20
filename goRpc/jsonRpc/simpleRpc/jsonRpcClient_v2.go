package main

import (
	"GoNet/goRpc/jsonRpc/simpleRpc/rpcObjects"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, e := net.Dial("tcp","localhost:1234")
	if e != nil {
		log.Fatalln("Error dialing: ", e)
	}

	cliCodec := jsonrpc.NewClientCodec(conn) // 建立针对客户端的json编解码器
	cli := rpc.NewClientWithCodec(cliCodec)

	args := &rpcObjects.Args{N: 7, M: 8}
	var reply int

	//e = conn.Call("Args.Multiply", args, &reply)
	//if e != nil {
	//	log.Fatalln("conn calling error: ", e)
	//	return
	//}
	//fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	// 异步调用
	call := cli.Go("Args.Multiply", args, &reply, nil)
	replyCall := <- call.Done
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
	fmt.Printf("replyCall: %v\n", replyCall)
}