package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// s := getCode("test@test.com")
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}
	if req.Method == http.MethodPost {
		e := req.FormValue("em")
		c.Value = e + `|` + getCode(e)
	}
	http.SetCookie(res, c)

	// 注意 這邊` + c.Value + `
	io.WriteString(res,
		`<!DOCTYPE html>
			<html>
				<body>
					<form method="Post">
						<input type="email" name="em">
						<input type="submit">
					</form>
					<a href="/authenticate">Validate this + `+c.Value+` + </a> 
				</body>
			</html>
		`)
}

func auth(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	//記得Redirect()要加上return
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if c.Value == "" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	s_list := strings.Split(c.Value, "|")
	email := s_list[0]
	mac := s_list[1]
	mac2 := getCode(email)

	if mac != mac2 {
		fmt.Println(`HMAC does not match`)
		fmt.Println(mac)
		fmt.Println(mac2)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	io.WriteString(res, `
	<!DOCTYPE html>
	<html>
		<body>
		<h1>`+mac+` - RECEIVED </h1>
		<h1>`+mac2+` - RECALCULATED </h1>
		</body>
	</html>
	`)
}

func getCode(m string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, m)
	// return string(h.Sum(nil)) //也可以，但這樣會亂碼
	return fmt.Sprintf("%x", h.Sum(nil))
}
