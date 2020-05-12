// 访问并读取页面内容
package main

import (
	"GoNet/goWeb"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("httpHead:---------------------")
	httpHead()
	fmt.Println("httpGet:---------------------")
	httpGet()
	fmt.Println("httpStatus:---------------------")
	httpStatus()
}

//向指定的URL发送一个简单的 http.Head(), 请求查看返回值
func httpHead(){
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.learnku.com/",
		"https://juejin.im/",
	}

	for _, url := range urls{
		resp, err := http.Head(url)
		goWeb.ErrorHandle(err, "httpHead :")
		fmt.Println(url, ": ", resp.Status)
	}
}

//向指定的URL发送一个简单的 http.Get(), 获取网页内容
func httpGet(){
	var url = "http://www.baidu.com/"
	resp, err := http.Get(url)
	goWeb.ErrorHandle(err, "httpGet Get:")
	data, err := ioutil.ReadAll(resp.Body)
	goWeb.ErrorHandle(err, "httpGet ReadAll:")
	fmt.Printf("Got: %q", string(data))
}

//向指定的URL发送一个简单的 http.Get(), 获取网页内容
func httpStatus(){
	type Status struct {
		Text string
	}

	type User struct {
		XMLName xml.Name
		Status  Status
	}
	// 发起请求查询推特Goodland用户的状态
	resp, err := http.Get("http://twitter.com/users/Googland.xml")
	goWeb.ErrorHandle(err, "httpStatus Get:")

	// 初始化XML响应的结构
	user := User{XMLName: xml.Name{Local: "user"}, Status: Status{""}}

	// 将XML解码到我们的结构中
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	goWeb.ErrorHandle(err, "httpStatus ReadAll:")
	fmt.Printf("body : %s ---", body)

	err = xml.Unmarshal(body, &user)
	goWeb.ErrorHandle(err, "httpStatus Unmarshal:")

	fmt.Printf("name: %s ", user.XMLName)
	fmt.Printf("status: %s", user.Status.Text)
}
