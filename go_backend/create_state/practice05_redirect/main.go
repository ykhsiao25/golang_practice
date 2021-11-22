package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	//流程 Get /barred HTTP/1.1 -> HTTP/1.1 200 OK -> Client填資料 -> POST /bar HTTP/1.1 -> HTTP/1.1 303 See Other
	// Get / HTTP/1.1 -> HTTP/1.1 200 OK
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)

	//以下兩行，可以用 http.Redirect(res, req, "/", http.StatusSeeOther)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusSeeOther)
}
func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
