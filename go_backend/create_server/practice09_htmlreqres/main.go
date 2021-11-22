package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl1 *template.Template

type hotdog int

//現在只是把 tpl.ExecuteTemplate() 從 直接func main()執行 移動到ServeHTTP() 而已(這樣才能讓web處理req and res)
//記得把 req.Form (or any req.DATA) 放到 ExecuteTemplate()內
// 每次ServeHTTP() 會執行兩次 (第二次的req.Form是空的 暫時不知道原因)
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//For all requests, ParseForm parses the raw query from the URL and updates r.Form.
	//parse the request, and store the data into req.Form
	err := req.ParseForm() //這行就會取得 fname的var value了
	fmt.Println(req.Form)
	if err != nil {
		log.Fatal(err)
	}

	tpl1.ExecuteTemplate(res, "index.gohtml", req.Form)
	fmt.Println("GG")
}
func init() {
	tpl1 = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
