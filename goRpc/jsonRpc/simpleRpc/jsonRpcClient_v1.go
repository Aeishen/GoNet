package main

import (
	"GoNet/goRpc/jsonRpc/simpleRpc/rpcObjects"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	conn, e := jsonrpc.Dial("tcp","localhost:1234")
	if e != nil {
		log.Fatalln("Error dialing: ", e)
	}

	args := &rpcObjects.Args{N: 7, M: 8}
	var reply int

	//e = conn.Call("Args.Multiply", args, &reply)
	//if e != nil {
	//	log.Fatalln("conn calling error: ", e)
	//	return
	//}
	//fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	// 异步调用
	call := conn.Go("Args.Multiply", args, &reply, nil)
	replyCall := <- call.Done
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
	fmt.Printf("replyCall: %v\n", replyCall)
}
