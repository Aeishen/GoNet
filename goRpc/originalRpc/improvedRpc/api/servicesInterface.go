/*
将rpc接口分为三个部分:
	①:服务的名字
	②:服务要实现的详细的方法列表
	③:注册该类型服务的函数
*/


package api

import (
	"GoNet/goRpc/originalRpc/improvedRpc/rpcObjects"
	"log"
	"net/rpc"
)

const ArgsServiceName = "ArgsServer"

type ArgsServiceInterface = interface {
	Multiply(req *rpcObjects.Args, resp *int) error
	Add(req *rpcObjects.Args, resp *int) error
}

// 注册服务时, 要求传入的对象满足 ArgsServiceInterface 接口
func RegisterArgsService(svc ArgsServiceInterface)  {
	err := rpc.RegisterName(ArgsServiceName, svc)
	if err != nil {
		log.Fatalln("rpc RegisterName error", err)
	}
}

//-------------------对客户端进行简单包装-----------------


//包装客户端类型, 该类型满足ArgsServiceInterface接口,这样客户端用户就可以直接通过接口对应的方法调用rpc函数
type ArgsServiceClient struct {
	*rpc.Client
}

var _ ArgsServiceInterface = (*ArgsServiceClient) (nil)


func (a *ArgsServiceClient) Multiply(req *rpcObjects.Args, resp *int) error {
	return a.Client.Call(ArgsServiceName + ".Multiply",req,resp)
}

func (a *ArgsServiceClient) Add(req *rpcObjects.Args, resp *int) error {
	return a.Client.Call(ArgsServiceName + ".Add",req,resp)
}

//包装拨号服务
func DailArgsService(network, address string)(*ArgsServiceClient, error)  {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &ArgsServiceClient{c}, nil
}
