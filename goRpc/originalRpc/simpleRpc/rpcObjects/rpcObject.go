// 定义一个"参数"类型, 该类型有一个Multiply方法

package rpcObjects

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}