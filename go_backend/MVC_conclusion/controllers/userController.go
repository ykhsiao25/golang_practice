package controllers

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/ykhsiao25/golang_practice/go_backend/MVC_conclusion/models"
	"github.com/ykhsiao25/golang_practice/go_backend/MVC_conclusion/session"
	"golang.org/x/crypto/bcrypt"
)

// time邏輯: 每存取一次cookie(session)，就要更新一下cookie.MaxAge成Length(除非要delete為-1)
// 每存取一次session就要更新一次sess.LastActivity = time.Now()，代表他有在使用
// (因為signup and login)都是直接新增sess到dbSessions(他就是變相update LastActivity了)
func (c Controller) Signup(res http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var user1 models.User
	if req.Method == http.MethodPost {
		//create the session
		ck, err := req.Cookie("session")
		if err != nil {
			uid := uuid.NewV4()
			ck = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
		}
		ck.MaxAge = session.Length
		http.SetCookie(res, ck)

		//parse the form
		username := req.FormValue("username")
		pwd := []byte(req.FormValue("password"))
		first := req.FormValue("first")
		last := req.FormValue("last")
		role := req.FormValue("role")

		//check the username
		if _, ok := session.DbUsers[username]; ok {
			http.Error(res, "Username already taken!", http.StatusForbidden)
			return
		}

		//encrypt the password
		d_pwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}

		//write to DB
		user1 = models.User{username, d_pwd, first, last, role}
		session.DbSessions[ck.Value] = models.Session{UserName: username, LastActivity: time.Now()}
		session.DbUsers[username] = user1
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
	c.tpl.ExecuteTemplate(res, "signup.html", user1)
}

func (c Controller) Login(res http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var user1 models.User
	if req.Method == http.MethodPost {
		//create the session
		ck, err := req.Cookie("session")
		if err != nil {
			uid := uuid.NewV4()
			ck = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
		}
		ck.MaxAge = session.Length
		http.SetCookie(res, ck)

		//parse the form
		username := req.FormValue("username")
		pwd := []byte(req.FormValue("password"))

		//check the username
		if _, ok := session.DbUsers[username]; !ok {
			http.Error(res, "Useranme do not match!", http.StatusForbidden)
			return
		}

		//compare the pasword
		user1 = session.DbUsers[username]
		err = bcrypt.CompareHashAndPassword(user1.Password, pwd)
		if err != nil {
			http.Error(res, "Wrong Password !!", http.StatusForbidden)
			return
		}
		//update the DB
		session.DbSessions[ck.Value] = models.Session{username, time.Now()}

		// Redirect the page
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(res, "login.html", user1)
}

func (c Controller) Logout(res http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	//get the cookie
	ck, _ := req.Cookie("session")

	//delete the session
	delete(session.DbSessions, ck.Value)

	//delete the cookie
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, ck)

	//clean all the sessions
	if time.Since(session.LastCleaned) > (time.Second * 30) {
		go session.Clean()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
