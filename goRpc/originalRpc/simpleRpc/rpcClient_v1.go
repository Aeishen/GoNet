package main

import (
	"./rpcObjects"
	"fmt"
	"log"
	"net/rpc"
)

func main()  {

	// DialHTTP在默认的HTTP RPC路径上侦听指定网络地址处的HTTP RPC服务器
	cli, err := rpc.DialHTTP("tcp","localhost:1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	args := &rpcObjects.Args{N: 7, M: 8}
	var reply int

	// 这个调用是同步的，所以需要等待结果返回
	err = cli.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	//异步调用
	//call1 := cli.Go("Args.Multiply", args, &reply, nil)
	//replyCall := <- call1.Done
	//fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
	//fmt.Printf("replyCall: %v\n", replyCall)
}