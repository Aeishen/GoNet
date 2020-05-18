package main

import (
	"./rpcObjects"
	"fmt"
	"log"
	"net/rpc"
)

func main()  {

	// DialHTTP在默认的HTTP RPC路径上侦听指定网络地址处的HTTP RPC服务器
	conn, err := rpc.DialHTTP("tcp","localhost:1234")
	if err != nil {
		log.Fatalln("Error dialing:", err)
	}

	args := &rpcObjects.Args{N: 7, M: 8}
	var reply int

	// 这个调用是同步的，所以需要等待结果返回
	err = conn.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatalln("conn calling error", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	//异步调用
	//Go方法是异步的，它返回一个 Call指针对象，它的Done是一个channel，如果服务返回，Done就可以得到返回的对象(实际是Call对象，包含Reply和error信息)
	//call := conn.Go("Args.Multiply", args, &reply, nil)
	//replyCall := <- call.Done
	//if replyCall.Error != nil {
	//	log.Fatalf("failed to call: %v", replyCall.Error)
	//} else {
	//	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
	//}
}