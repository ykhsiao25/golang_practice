package main

import (
	"html/template"
	"net/http"

	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/MVC_conclusion/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController()
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}
