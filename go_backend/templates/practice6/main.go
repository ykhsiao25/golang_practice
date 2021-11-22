package main

import (
	"log"
	"os"
	"text/template"
)

var tpl1 *template.Template

//撰寫template順序應該是這樣
//寫html template，設計好最高到最低的struct{}及其對應的func()要執行出什麼，有個框架
//再到go把每一層的struct{} 都宣告好
//以及每一層用到的func()都寫進template.funcMap{}內
//然後實作每一層的func()
//最後 New() + Funcs() + parse() instantiate  要執行的template 再執行tpl()並帶入其需要用的data instance
func init() {
	tpl1 = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	err1 := tpl1.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}
