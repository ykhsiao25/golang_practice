package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// client 打 "/" -> redirect to /signup -> render the html -> client POST the form ->
// make the cookie,  get the form value and store the userinfo into db ->
// redirect to the index page -> render the index html -> click the bar page ->
// render the bar html
type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

// cookie name 和 username是啥沒有關係，cookie.Name只是要拿出那個cookie而已
var tpl *template.Template
var dbSessions = make(map[string]string)
var dbUsers = make(map[string]user)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

// index, bar, signup func()
// if no signup, then signup
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
func signup(res http.ResponseWriter, req *http.Request) {
	//if logged in, redirect to the 首頁
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var user1 user
	//get the form value
	if req.Method == http.MethodPost {
		// get the cookie
		c, err := req.Cookie("session")
		if err != nil {
			uid := uuid.NewV4()
			c = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
			http.SetCookie(res, c)
		}

		u_name := req.FormValue("username")
		pwd := req.FormValue("password")
		f_name := req.FormValue("firstname")
		l_name := req.FormValue("lastname")

		//check username exists or not(用userinfo table去找裡面有沒有user了)
		if _, ok := dbUsers[u_name]; ok {
			// 403 = StatusForbidden = serve knows the request, but refuse to approve it
			http.Redirect(res, req, "/", http.StatusForbidden)
			return
		}

		user1 = user{u_name, pwd, f_name, l_name}

		dbSessions[c.Value] = u_name
		dbUsers[u_name] = user1

		//after signup, redirect to the index page
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", user1)
}
