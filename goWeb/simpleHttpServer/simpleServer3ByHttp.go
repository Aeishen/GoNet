//网页服务器2，可以让用户输入一连串的数字，然后将它们打印出来，计算出这些数字的均值和中值

package main

import (
	"GoNet/goWeb"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const form3 = `<html><body><form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></html></body>`

const myError = `<p class="error">%s</p>`

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

var pageTop = ""
var pageBottom = ""

func CalServer(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-type","text/html")

	// ParseForm都会从URL解析原始查询并更新req.Form。
	err := req.ParseForm()
	_, _ = fmt.Fprint(w, pageTop, form3)
	if err != nil{
		_, _ = fmt.Fprintf(w, myError, err)
	}else{
		if numbers, message, ok := processRequest(req); ok {
			stats := getStats(numbers)
			_, _ = fmt.Fprint(w, formatStats(stats))
		} else if message != "" {
			_, _ = fmt.Fprintf(w, myError, message)
		}
	}
	_, _ = fmt.Fprint(w, pageBottom)
}


// 处理请求
func processRequest(req *http.Request) ([]float64, string, bool) {
	var numbers []float64
	var text string
	if slice, found := req.Form["numbers"]; found && len(slice) > 0 {
		//处理如果网页中输入的是中文逗号
		if strings.Contains(slice[0], "&#65292") {
			text = strings.Replace(slice[0], "&#65292;", " ", -1)
		} else {
			text = strings.Replace(slice[0], ",", " ", -1)
		}

		//将字符串s围绕一个或多个连续的空白字符的每个实例进行拆分，返回s的子字符串切片
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}

func main() {
	http.HandleFunc("/", CalServer)
	err := http.ListenAndServe("localhost:8080", nil)
	goWeb.ErrorHandle(err, "ListenAndServe :")
}