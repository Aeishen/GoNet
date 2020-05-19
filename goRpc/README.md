# goRpc

###什么是RPC
在分布式计算，远程过程调用（英语：Remote Procedure Call，缩写为 RPC）是一个计算机通信协议。
该协议允许运行于一台计算机的程序调用另一个地址空间（通常为一个开放网络的一台计算机）的子程序，
而程序员就像调用本地程序一样，无需额外地为这个交互作用编程（无需关注细节）。
RPC是一种服务器-客户端（Client/Server）模式，经典实现是一个通过发送请求-接受回应进行信息交互的系统。
如果涉及的软件采用面向对象编程，那么远程过程调用亦可称作远程调用或远程方法调用

RPC是一种进程间通信的模式，程序分布在不同的地址空间里。如果在同一主机里，RPC可以通过不同的虚拟地址空间
（即便使用相同的物理地址）进行通讯，而在不同的主机间，则通过不同的物理进行交互。许多技术（常常是不兼容）
都是基于这种概念而实现的。

用通俗易懂的语言描述就是：RPC允许跨机器、跨语言调用计算机程序方法。打个比方，
我用go语言写了个获取用户信息的方法getUserInfo，并把go程序部署在阿里云服务器上面，
现在我有一个部署在腾讯云上面的php项目，需要调用golang的getUserInfo方法获取用户信息，
php跨机器调用go方法的过程就是RPC调用。

###RPC流程
1. 客户端调用客户端stub（client stub）。这个调用是在本地，并将调用参数push到栈（stack）中。
2. 客户端stub（client stub）将这些参数包装，并通过系统调用发送到服务端机器。打包的过程叫 marshalling。（常见方式：XML、JSON、二进制编码）
3. 客户端本地操作系统发送信息至服务器。（可通过自定义TCP协议或HTTP传输）
4. 服务器系统将信息传送至服务端stub（server stub）。
5. 服务端stub（server stub）解析信息。该过程叫 unmarshalling。
6. 服务端stub（server stub）调用程序，并通过类似的方式返回给客户端。

###golang中如何实现RPC
在golang中实现RPC非常简单，有封装好的官方库和一些第三方库提供支持。Go RPC可以利用tcp或http来传递数据，
可以对要传递的数据使用多种类型的编解码方式。golang官方的net/rpc库使用encoding/gob进行编解码，支持tcp
或http数据传输方式，由于其他语言不支持gob编解码方式，所以使用net/rpc库实现的RPC方法没办法进行跨语言调用。

golang官方还提供了net/rpc/jsonrpc库实现RPC方法，JSON RPC采用JSON进行数据编解码，因而支持跨语言调用。
但目前的jsonrpc库是基于tcp协议实现的，暂时不支持使用http进行数据传输。

除了golang官方提供的rpc库，还有许多第三方库为在golang中实现RPC提供支持，大部分第三方rpc库的实现都是使用
protobuf进行数据编解码，根据protobuf声明文件自动生成rpc方法定义与服务注册代码，在golang中可以很方便的进
行rpc服务调用。

###golang中RPC框架的两个特色设计
1. rpc数据打包时可以通过插件实现自定义的编码和解码
2. rpc建立在抽象的io.ReadWriteCloser接口之上, 我们可以将rpc架设在不同的通信协议上


