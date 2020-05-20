package main

import (
	"GoNet/goRpc/originalRpc/improvedRpc/api"
	"GoNet/goRpc/originalRpc/improvedRpc/rpcObjects"
	"fmt"
	"log"
)



func main() {
	conn, err := api.DailArgsService("tcp", "localhost:1234")
	if err != nil {
		log.Fatalln("Error dialing:", err)
	}

	args := &rpcObjects.Args{N: 7, M:8}
	var reply int

	err = conn.Multiply(args, &reply)
	if err != nil {
		log.Fatalln("conn calling Multiply error", err)
	}
	fmt.Printf("Multiply: %d * %d = %d\n", args.N, args.M, reply)

	err = conn.Add(args, &reply)
	if err != nil {
		log.Fatalln("conn calling Add error", err)
	}
	fmt.Printf("Add: %d + %d = %d\n", args.N, args.M, reply)
}
