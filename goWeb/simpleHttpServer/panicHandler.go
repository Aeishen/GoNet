// 使用闭包的方法处理错误, 使得程序更加健壮

package main

import (
	"log"
	"net/http"
)

type HandleFnc func(w http.ResponseWriter, r *http.Request)

func LogPanics(fun HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil{
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)

				//默认出现 panic 只会记录日志，页面就是一个无任何输出的白页面， 可以给页面一个错误信息，如下面的示例返回了一个 500
				http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		fun(w, r)
	}
}
