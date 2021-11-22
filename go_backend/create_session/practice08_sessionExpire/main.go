package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}
type session struct {
	userid       string
	lastActivity time.Time
}

var tpl *template.Template
var dbSessions = map[string]session{}
var dbUsers = map[string]user{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

//這邊只會執行一次
func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	dbSessionsCleaned = time.Now()
	fmt.Println("init()")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	user1 := getUser(res, req)
	showSessions()
	tpl.ExecuteTemplate(res, "index.gohtml", user1)
}

func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	user1 := getUser(res, req)
	if user1.Role != "007" {
		http.Error(res, "You have to be 007!", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "bar.gohtml", user1)
}
func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var user1 user
	if req.Method == http.MethodPost {
		//create the session(只要有session，就要調整cookie.MaxAge)
		c, err1 := req.Cookie("session")
		if err1 != nil {
			uid := uuid.NewV4()
			c = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
		}
		c.MaxAge = sessionLength
		//因為把MaxAge重新set了，所以要把SetCookie()寫在外面
		http.SetCookie(res, c)

		//parse the form
		u_name := req.FormValue("username")
		pwd := req.FormValue("password")
		b_pwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		f_name := req.FormValue("firstname")
		l_name := req.FormValue("lastname")
		r := req.FormValue("role")
		user1 = user{u_name, b_pwd, f_name, l_name, r}

		//examine the username
		sess := dbSessions[c.Value]
		if _, ok := dbUsers[sess.userid]; ok {
			http.Error(res, "Username was used!!", http.StatusForbidden)
			return
		}

		//write to db
		dbSessions[c.Value] = session{u_name, time.Now()}
		dbUsers[u_name] = user1
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", user1)
}

//個人認為login應該只要check username and password OK 即可 (透過user.password和FormValue()比較)
//但這邊考慮可能本身server和client沒有session cookie
func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		//create the session
		//dbUsers[username] 本來就有資訊，但dbSessions[uid]可能沒有，所以要新增
		c, err2 := req.Cookie("session")
		if err2 != nil {
			uid := uuid.NewV4()
			c = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)

		//parse the form
		u_name := req.FormValue("username")
		pwd := req.FormValue("password")

		//check the username(看存不存在而已，沒辦法比對)
		user1, ok := dbUsers[u_name]
		if !ok {
			http.Error(res, "No username here", http.StatusForbidden)
			return
		}
		//check the password
		err1 := bcrypt.CompareHashAndPassword(user1.Password, []byte(pwd))
		if err1 != nil {
			http.Error(res, "Wrong Password here", http.StatusForbidden)
			return
		}
		dbSessions[c.Value] = session{u_name, time.Now()}

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

//logout是把dbSessions[uid]刪掉(delete 這個session)
func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	c, err1 := req.Cookie("session")
	if err1 != nil {
		http.Error(res, "No user should sign out !", http.StatusForbidden)
		return
	}
	delete(dbSessions, c.Value)

	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(res, c)
	fmt.Println("session time out 1")

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		fmt.Println("session time out 2")
		go cleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
