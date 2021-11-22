package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl1 *template.Template
var tpl2 *template.Template
var tpl3 *template.Template
var tpl4 *template.Template
var tpl5 *template.Template
var tpl6 *template.Template
var tpl7 *template.Template
var tpl8 *template.Template
var tpl9 *template.Template
var tpl10 *template.Template
var tpl11 *template.Template
var tpl12 *template.Template

//撰寫template順序應該是這樣
//寫html template，設計好最高到最低的struct{}及其對應的func()要執行出什麼，有個框架
//再到go把每一層的struct{} 都宣告好
//以及每一層用到的func()都寫進template.funcMap{}內
//然後實作每一層的func()
//最後 New() + Funcs() + parse() instantiate  要執行的template 再執行tpl()並帶入其需要用的data instance

//passing function into template to process the data
//special structure
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": func(s string) string {
		st := strings.TrimSpace(s)
		if len(st) >= 3 {
			return st[:3]
		}
		return st
	},
}
var fm2 = template.FuncMap{
	"fdateMDY": func(t time.Time) string {
		//這是時間的template，且必須是這個時間才對(據說是golang誕生日)
		//MM-DD-YYYY
		return t.Format("01-02-2006")
	},
}
var fm3 = template.FuncMap{
	"fdbl": func(x int) int {
		return x + x
	},
	"fsql": func(x int) float64 {
		return math.Pow(float64(x), 2)
	},
	"fsqrt": func(x float64) float64 {
		return math.Sqrt(x)
	},
}

func init() {
	tpl1 = template.Must(template.ParseFiles("index.html"))
	tpl2 = template.Must(template.ParseFiles("index2.html"))
	tpl3 = template.Must(template.ParseFiles("index3.html"))
	tpl4 = template.Must(template.ParseFiles("index4.html"))
	tpl5 = template.Must(template.ParseFiles("index5.html"))
	tpl6 = template.Must(template.ParseFiles("index6.html"))

	//function passing
	// Since the templates created by ParseFiles are named by the base names of the argument files,
	// t should usually have the name of one of the (base) names of the files.
	// https: //pkg.go.dev/text/template#Template.ParseFiles
	//要嘛 New("filename") + Execute(writer, data) 要嘛就 New("") + ExecuteTemplate(writer, filename, data)
	// create template(with name) -> attach funcMap -> parseFiles -> Execute() -> check error
	tpl7 = template.Must(template.New("index7.html").Funcs(fm).ParseFiles("index7.html"))
	tpl8 = template.Must(template.New("index_time.html").Funcs(fm2).ParseFiles("index_time.html"))
	tpl9 = template.Must(template.New("index_pipeline.html").Funcs(fm3).ParseFiles("index_pipeline.html"))

	//global func() (html)
	tpl10 = template.Must(template.ParseFiles("index_global_func_index.html"))
	tpl11 = template.Must(template.ParseFiles("index_global_func_and.html"))
	//要用到別的 file 的template，要記得在這邊parse()，不然找不到
	tpl12 = template.Must(template.ParseFiles("index_nested_template.html", "nested_template.gohtml"))
}

func main() {
	//slice
	sages := []string{"Gandhi", "MLK", "Buddha"}
	// err := tpl1.Execute(os.Stdout, sages)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err2 := tpl2.Execute(os.Stdout, sages)
	if err2 != nil {
		log.Fatal(err2)
	}

	//map
	m1 := map[string]string{
		"india":    "Gandhi",
		"US":       "MLK",
		"Meditate": "Buddha",
	}
	err3 := tpl3.Execute(os.Stdout, m1)
	if err3 != nil {
		log.Fatal(err3)
	}

	type sage struct {
		Name  string
		Motto string
	}
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	err4 := tpl4.Execute(os.Stdout, buddha)
	if err4 != nil {
		log.Fatal(err4)
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}
	sage_slice := []sage{gandhi, jesus}
	err5 := tpl5.Execute(os.Stdout, sage_slice)
	if err5 != nil {
		log.Fatal(err5)
	}

	type car struct {
		Manufacter string
		Model      string
		Doors      string
	}

	type items struct {
		Wisdom    []sage
		Transport []car
	}

	f := car{
		Manufacter: "A",
		Model:      "m1",
		Doors:      "single",
	}
	g := car{
		Manufacter: "B",
		Model:      "m2",
		Doors:      "double",
	}

	its := items{
		Wisdom:    sage_slice,
		Transport: []car{f, g},
	}
	err6 := tpl6.Execute(os.Stdout, its)
	if err6 != nil {
		log.Fatal(err6)
	}

	//可以直接宣告(如果只用一次的話)
	//注意struct 的 property and method不要用逗號
	its_refactor := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sage_slice,
		Transport: []car{f, g},
	}
	err6_2 := tpl6.Execute(os.Stdout, its_refactor)
	if err6_2 != nil {
		log.Fatal(err6_2)
	}

	// err7 := tpl7.ExecuteTemplate(os.Stdout, "index7.html", its_refactor)
	err7 := tpl7.Execute(os.Stdout, its_refactor)
	if err7 != nil {
		log.Fatal(err7)
	}

	//data 代表想pass進去東西是什麼 也就是 {{.}}你想放什麼
	err8 := tpl8.Execute(os.Stdout, time.Now())
	if err8 != nil {
		log.Fatalln(err8)
	}

	err9 := tpl9.Execute(os.Stdout, 3)
	if err9 != nil {
		log.Fatalln(err9)
	}

	s_slice := []string{"zero", "one", "two", "three"}
	tpl10_s := struct {
		Words []string
		Lname string
	}{
		Words: s_slice,
		Lname: "Jeff",
	}
	err10 := tpl10.Execute(os.Stdout, tpl10_s)
	if err10 != nil {
		log.Fatal((err10))
	}

	type user struct {
		Name  string
		Motto string
		Admin bool
	}

	u1 := user{
		Name:  "A",
		Motto: "a",
		Admin: false,
	}

	u2 := user{
		Name:  "B",
		Motto: "b",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}
	users := []user{u1, u2, u3}
	err11 := tpl11.Execute(os.Stdout, users)
	if err11 != nil {
		log.Fatal(err11)
	}

	err12 := tpl12.Execute(os.Stdout, 42)
	if err12 != nil {
		log.Fatal(err12)
	}
}
