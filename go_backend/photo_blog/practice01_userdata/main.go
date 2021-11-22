package main

import (
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", c.Value)
}
func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		uid := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(res, c)
	}
	return c
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
