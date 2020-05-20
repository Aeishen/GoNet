//采用http协议作为rpc载体

package main

import (
	"../../originalRpc/simpleRpc/rpcObjects"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	_ = rpc.RegisterName("Args", new(rpcObjects.Args))        // 注册rpc服务

	http.HandleFunc("/httpRpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			w,
			r.Body,
		}

		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	_ = http.ListenAndServe("localhost:1234", nil)
}
