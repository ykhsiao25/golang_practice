package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/signup", http.StatusSeeOther)
		return
	}
	user1 := getUser(req)
	tpl.ExecuteTemplate(res, "index.gohtml", user1)
}

func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/signup", http.StatusSeeOther)
		return
	}
	user1 := getUser(req)
	tpl.ExecuteTemplate(res, "bar.gohtml", user1)
}

//remember to redirect to the index page after parsing the form values
// (Don't need to render the .gohtml to the browser)
func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var user1 user
	if req.Method == http.MethodPost {
		c, err1 := req.Cookie("session")
		if err1 != nil {
			uid := uuid.NewV4()
			fmt.Println("Create a cookie!")
			c = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
				Path:  "/",
			}
			http.SetCookie(res, c)
		}
		u_name := req.FormValue("username")
		passwd := req.FormValue("password")
		f_name := req.FormValue("firstname")
		l_name := req.FormValue("lastname")

		if _, ok := dbUsers[u_name]; ok {
			http.Error(res, "UserName was already taken.", http.StatusForbidden)
			return
		}

		b_passwd, err2 := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
		if err2 != nil {
			http.Error(res, "Some Interal Server Error", http.StatusInternalServerError)
			return
		}
		user1 = user{u_name, b_passwd, f_name, l_name}

		dbSessions[c.Value] = u_name
		dbUsers[u_name] = user1

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", user1)
}
