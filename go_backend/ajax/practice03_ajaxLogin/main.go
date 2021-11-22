package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type session struct {
	Username     string
	lastActivity time.Time
}
type user struct {
	Username string
	Password []byte
	First    string
	Last     string
	Role     string
}

var tpl *template.Template
var dbSessions = map[string]session{}
var dbUsers = map[string]user{}
var dbSessionsCleaned time.Time //session需要被delete的時間

const sessionLength int = 30 //session允許存活之時長

//這邊只會執行一次
func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
	dbSessionsCleaned = time.Now()
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	// added new route
	http.HandleFunc("/checkUserName", checkUserName)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// 首頁: 看登入及註冊選項
func index(res http.ResponseWriter, req *http.Request) {
	user1 := getUser(res, req)
	showSessions()
	tpl.ExecuteTemplate(res, "index.html", user1)
}

//限制權限
// bar(看符合身分否?) == (!alreadyLoggedIn?) + (getUser.Role == ?)
func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	user1 := getUser(res, req)
	if user1.Role != "007" {
		http.Error(res, "You have to be 007 !", http.StatusForbidden)
	}
	showSessions()
	tpl.ExecuteTemplate(res, "bar.html", user1)
}

// signup == (alreadyLoggedIn?) + (create session) + (parse the form) +
//(examine the username) + (encrypt the pwd) + (write to db)
func signup(res http.ResponseWriter, req *http.Request) {
	// (alreadyLoggedIn?)
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var user1 user
	//create the session(只要有session，就要調整cookie.MaxAge)
	if req.Method == http.MethodPost {
		// (create session)
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

		//(parse the form)
		u_name := req.FormValue("username")
		pwd := []byte(req.FormValue("password"))
		f_name := req.FormValue("firstname")
		l_name := req.FormValue("lastname")
		role := req.FormValue("role")

		//examine the username
		_, ok := dbSessions[c.Value]
		if ok {
			http.Error(res, "username has been used!", http.StatusForbidden)
		}
		//記得加密
		d_pwd, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

		//(write to db)
		user1 = user{u_name, d_pwd, f_name, l_name, role}
		dbSessions[c.Value] = session{u_name, time.Now()}
		dbUsers[u_name] = user1
	}
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

//個人認為login應該只要check username and password OK 即可 (透過user.password和FormValue()比較)
//但這邊考慮可能本身server和client沒有session cookie
//login == (create a session) + (parse username + pwd) +
// (username exist ?) + (compare password) + (back to 首頁)
func login(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		//(create a session)
		c, err := req.Cookie("session")
		if err != nil {
			uid := uuid.NewV4()
			c = &http.Cookie{
				Name:  "session",
				Value: uid.String(),
			}
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		// (parse username + pwd)
		u_name := req.FormValue("username")
		pwd := []byte(req.FormValue("password"))

		//(username exist ?)
		user1, ok := dbUsers[u_name]
		if !ok {
			http.Error(res, "NO User Exist!", http.StatusForbidden)
			return
		}

		//(compare password)
		err1 := bcrypt.CompareHashAndPassword(user1.Password, pwd)
		if err1 != nil {
			http.Error(res, "NOT correct password", http.StatusForbidden)
			return
		}

		//dbUsers[username] 本來就有資訊，但dbSessions[uid]可能沒有，所以要新增
		dbSessions[c.Value] = session{u_name, time.Now()}

		//back to 首頁
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
}

// (!alreadyLoggedIn) + (get cookie) + (delete session) + (delete cookie) + (cleanup sessions)
//這邊strategy是在有人logout的時候，每個session也要檢查是不是30秒沒動了，就一起砍掉
func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")

	//delete dbSessions
	delete(dbSessions, c.Value)

	// delete the cookie (MaxAge < 0 means delete the cookie)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	fmt.Println("session time out 1")

	if time.Since(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

// 建一個sampleUsers，每次都直接讀取req.Body(我猜是因為AJAX，每次都只有一部分)，然後直接把string寫到response
// 之後會在browser端去判斷從server端接收到的value，再去做處理
func checkUserName(res http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"test@example.com": true,
		"jame@bond.com":    true,
		"moneyp@uk.gov":    true,
	}

	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	s := string(bs)
	fmt.Println("Username: ", s)
	// fmt.Println("sampleUsers[s]: ", sampleUsers[s])
	//注意 這邊千萬不能用Fprintln()!!!!! 否則後端會撈不到data(他會多一個'\n')
	fmt.Fprint(res, sampleUsers[s])
}
