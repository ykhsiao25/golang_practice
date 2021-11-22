package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl1 *template.Template

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method string
		URL    *url.URL
		// Submissions url.Values
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl1.ExecuteTemplate(res, "index.gohtml", data)
}

func init() {
	tpl1 = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
