// 定义一个"参数"类型, 该类型有一个Multiply方法

package rpcObjects

import (
	"GoNet/goRpc/protoRpc/simpleRpc/mypb"
)

type Args struct {}

func (t *Args) Multiply(req *mypb.ArgsReq, resp *mypb.ArgsResp) error {
	resp.Reply = req.GetN() * req.GetN()
	return nil
}