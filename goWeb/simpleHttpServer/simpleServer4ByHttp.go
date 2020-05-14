package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// 使用expvar包, 创建一个变量, 一般用于服务器中的操作计数器, 此处helloRequests 是一个 int64 类型的变量
var helloRequests = expvar.NewInt("hello-requests")

//
var webroot = flag.String("root", "/home/user", "web root directory")
var booleanFlag = flag.Bool("boolean", true, "another flag for testing")

// 简单的服务器计数器，发布它将设置值
type Counter struct {
	n int
}

//
type Chan chan int


func main() {
	flag.Parse()

	// 当 url 中的地址不存在（没有对应的路由）时，就会去匹配 / 对应的处理函数（ Logger()），它会在页面中显示一个 oops ，
	// 并且在 header 中写入 404（可以通过浏览器调试模式的 console 或者 network 查看，直接是看不到 404 的），
	// 在命令行窗口（也可以理解成日志文件）中记录下错误信息
	http.Handle("/", http.HandlerFunc(Logger))

	http.Handle("/go/hello", http.HandlerFunc(HelloServer))

	ctr := new(Counter)

	// Counter 对象 ctr 有一个 String () 方法，所以它就实现了 Var 接口, ( 因为publish 的第二个参数是个 Var
	// 接口，所以想要发布的结构体必须实现这个接口), ctr计数器直接作为一个变量被发布
	expvar.Publish("counter", ctr)

	// ctr 实现了 ServeHTTP 方法，就实现了 Handler 接口
	http.Handle("/counter", ctr)

	// FileServer 返回一个 root 参数的值为根目录的文件来处理 HTTP 请求。通过 http.Dir 去使用操作系统的文件系统
	// 可以在 /tmp 目录下创建一个 ggg.html , 再访问 /go/ggg.html 的时候就会直接在浏览器中显示 ggg.html 的内容
	// http.Handle("/go/", http.FileServer(http.Dir("/tmp")))
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))

	// 通过 flag.VisitAll 函数去遍历所有的 flags （前面的两个命令行参数），打印他们的变量名、值、默认值（如果值不是默认值的时候）。
	http.Handle("/flags", http.HandlerFunc(FlagServer))

	// 遍历 os.Args 去打印所有的命令行参数；如果没有就只会打印程序的名称（可执行文件的目录,这个有点类似 Linux 的 $0 $1 $2 。
	http.Handle("/args", http.HandlerFunc(ArgServer))

    // 通道的 ServeHTTP 方法在每个新请求中显示来自通道的下一个整数。所以一个 Web 服务器可以从一个通道接收响应，由另一个函数填充（甚至是客户端）
	http.Handle("/chan", ChanCreate())


	http.Handle("/date", http.HandlerFunc(DateServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}


func Logger(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.String())
	w.WriteHeader(404)
	_, _ = w.Write([]byte("oops"))
}

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	helloRequests.Add(1) //操作计数器加1
	_, _ = io.WriteString(w, "hello, world!\n")
}

func (ctr *Counter) String() string { return fmt.Sprintf("%d", ctr.n) }

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // n 自增
		ctr.n++
	case "POST": // set n to posted value
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			_, _ = fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			_, _ = fmt.Fprint(w, "counter reset\n")
		}
	}
	_, _ = fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprint(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			_, _ = fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			_, _ = fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

func ArgServer(w http.ResponseWriter, _ *http.Request) {
	for _, s := range os.Args {
		_, _ = fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}

// 重定向输出
func DateServer(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r, w, err := os.Pipe()  //返回一对连接的文件
	if err != nil {
		_, _ = fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}

	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	_ = w.Close()
	if err != nil {
		_, _ = fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	_, _ = io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		_, _ = fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		_, _ = fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}