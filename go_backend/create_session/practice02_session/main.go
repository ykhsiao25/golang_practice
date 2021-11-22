package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
	LoggedIn bool
}

//流程是這樣的
// client 打 "/" 來 server -> server 幫client建立cookie + 回 index.gohtml response ->
// save the cookie into the client local storage ->
// (Loop)
// client submit the index.gohtml form ->
// browser parses the form, posts the form and sends the cookie to the server ->
// server gets the cookie + gets the userinfo(maybe empty string) + gets the form value and stores userinfo into map(db) ->
// reponse back to the client
// (Loop ends)

var tpl *template.Template

//不可以寫成var dbSessions map[string]string(因為這只是宣告，本身為nil)
// var dbSessions = make(map[string]string)
var dbSessions = map[string]string{}

var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

//建立一個table (map[uuid string]userid string)
//再利用這個table去拿user info(map[userid string]user userinfo)
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	//get the cookie, if no cookie, then new a cookie
	//if you have a cookie, then you must have an uuid
	c, err := req.Cookie("session")
	if err != nil {
		uid := uuid.NewV4()
		//記得要加上 "&"
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
			Path:  "/",
		}
		http.SetCookie(res, c)
	}

	//if the user already exists, get the userinfo
	var user1 user
	//idiom ok, if no ok variable to get the result, then if error, panic
	// even 重新整理，依然不是nil or error，他會是golang預設的default value(ex: string "")
	if userid, ok := dbSessions[c.Value]; ok {
		user1 = dbUsers[userid]
	}

	//golang的語言特性，讓user{}不會是error
	//Process the submission form
	if req.Method == http.MethodPost {
		u_name := req.FormValue("username")
		f_name := req.FormValue("firstname")
		l_name := req.FormValue("lastname")
		login := req.FormValue("loggedIn") == "on"
		user1 = user{u_name, f_name, l_name, login}

		dbUsers[u_name] = user1
		dbSessions[c.Value] = u_name
	}
	//因為上面已經有宣告過user1 (var user1 user)，所以第一次(還沒有cookie)的response會user{"","","",false}(就default)
	//這邊沒有Redirect()，就在這邊一直重複執行(就是不斷把index.gohtml寫到res)
	tpl.ExecuteTemplate(res, "index.gohtml", user1)

}

// if you have a cookie, and it's our server's cookie, then get the bar.html
func bar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		//if no cookie,  redirect to the index page
		http.Redirect(res, req, "/", http.StatusSeeOther)
		//return 一定要寫，否則會error
		return
	}
	userid, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	user1 := dbUsers[userid]
	tpl.ExecuteTemplate(res, "bar.gohtml", user1)
}
