package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	// http.HandleFunc("/", index)
	http.HandleFunc("/", index2)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}
func index2(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index2.html", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	s := `Here is some text from foo`
	fmt.Fprintln(res, s)
}
