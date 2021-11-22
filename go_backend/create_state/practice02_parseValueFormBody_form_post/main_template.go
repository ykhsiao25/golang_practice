package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func foo(res http.ResponseWriter, req *http.Request) {
	//這就單純取得 FormValue()的API而已(就html 那個form)
	// http://127.0.0.1:8080/?q=12 這樣打(if method="GET")
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
