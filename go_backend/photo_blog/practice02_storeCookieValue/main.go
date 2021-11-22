package main

import (
	"net/http"
	"strings"
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
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	c = appendValue(res, c)
	s := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(res, "index.gohtml", s)
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

func appendValue(res http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Values
	p1 := "disneyland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"
	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	http.SetCookie(res, c)

	return c
}
