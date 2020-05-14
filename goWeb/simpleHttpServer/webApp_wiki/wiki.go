// 一个简单的web应用, 实现了一组页面的显示、编辑、和保存(使用命令行运行, 否则找不到文件路径)

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

// 一个wiki页面需要一个标题和文本内容, 内容是字节切片。
type Page struct {
	Title string
	Body []byte
}

type HandleFun func(http.ResponseWriter, *http.Request, string)

const lenPath = len("/view/")
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")  // 用来验证标题的输入是不是由字母与数字组成
var templates = make(map[string]*template.Template)            // 存储创建出来的 html 文件( 称为 模板缓存 )



func init() {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html")) // 将模板文件转换成一个 *template.Template 并存储
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// 封装处理函数
func makeHandler(fn HandleFun) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]

		//验证输入的标题
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

// 浏览处理(存在就直接展示，不存在就跳转到编辑)
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)

	// 文件不存在
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)  // 重定向到这个标题对应的编辑页面
		return
	}

	// 渲染模板并写入到 ResponseWriter w 中
	renderTemplate(w, "view", p)
}

// 编辑处理(存在就修改，不存在就添加)
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)

	// 文件不存在
	if err != nil {
		p = &Page{Title: title}
	}

	// 渲染模板并写入到 ResponseWriter w 中
	renderTemplate(w, "edit", p)
}

// 保存处理
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") //提取名字为 body 的文本域字段的内容
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // 重定向到这个标题对应的浏览页面
}


//使模板和结构体输出到页面(使用了模板去合并结构体与 html 模板中的数据)
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p) //调用一个模板，将 Page 结构体 p 作为一个参数在模板中进行替换，并且写入到 ResponseWriter w 中
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


//创建并保存指定标题的文本文件
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) 	// 文件仅具有当前用户的读写权限
}


// 加载指定标题的文本文件
func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) // 如果文件被找到则将它的内容读取到一个本地的字符串类型的 body 变量中
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}