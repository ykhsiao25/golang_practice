package controllers

import (
	"html/template"
	"net/http"

	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/MVC_conclusion/session"
)

type Controller struct {
	tpl *template.Template
}

func NewController(tpl *template.Template) *Controller {
	return &Controller{tpl}
}
func (c Controller) Index(res http.ResponseWriter, req *http.Request) {
	user1 := session.GetUser(req, req)
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(res, "index.html", user1)
}
func (c Controller) Bar(res http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	user1 := session.GetUser(res, req)
	if user1.Role != "007" {
		http.Error(res, "You need to be 007!!", http.StatusForbidden)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(res, "bar.html", user1)
}
