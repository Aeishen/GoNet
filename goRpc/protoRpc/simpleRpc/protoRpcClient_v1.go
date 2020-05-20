package main

import (
	"GoNet/goRpc/protoRpc/simpleRpc/mypb"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	conn, e := rpc.Dial("tcp","localhost:1234")
	if e != nil {
		log.Fatalln("Error dialing:", e)
	}

	args := &mypb.ArgsReq{N: 7, M: 8}
	r := &mypb.ArgsResp{}

	e = conn.Call("Args.Multiply",args,r)
	if e != nil {
		log.Fatalln("conn calling error", e)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, r.Reply)
}