# protoRpc
为了实现跨语言调用，在golang中实现RPC方法的时候我们应该选择一种跨语言的数据编解码方式，比如JSON，
前面的jsonrpc可以满足此要求，但是也存在一些缺点，比如不支持http传输，数据编解码性能不高等。于是呢，
一些第三方rpc库都选择采用protobuf进行数据编解码，并提供一些服务注册代码自动生成功能。这部分的例子
我们使用protobuf来定义RPC方法及其请求响应参数，并使用第三方的protorpc库来生成RPC服务注册代码。

###什么是protobuf
Google Protocol Buffer( 简称 Protobuf) 是 Google 公司内部的混合语言数据标准。Protocol Buffers 
是一种轻便高效的结构化数据存储格式，可以用于结构化数据串行化，或者说序列化。它很适合做数据存储或 RPC 
数据交换格式。可用于通讯协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。


###如何安装protobuf
自己百度
    
###如何使用protobuf
1. 在golang中安装protobuf相关的库: go get -u github.com/golang/protobuf/{protoc-gen-go,proto}
2. 编写.proto文件: .../xxx/xxx/user.proto

    ```syntax = "proto3";
     package pb;
     
     message user {
         int32 id = 1;
         string name = 2;
     }
     
     message multi_user {
         repeated user users = 1;
     }
3. 基于.proto文件生成数据操作代码: 在文件所在目录下执行，执行命令完成，在与user.proto文件同级的目录下生成了一个user.pb.go文件

    `protoc --go_out=. user.proto`