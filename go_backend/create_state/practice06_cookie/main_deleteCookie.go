package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//把接下來要顯示的page or content寫在這個func()
func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set">Set a cookie</a></h1>`)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	fmt.Fprintln(res, `<h1><a href="/read">Read the cookie</a></h1>`)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Println("No cookies can be read ! ")
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	//若要在html顯示要得value，則加上 "%v"
	fmt.Fprintf(res, `<h1>Your Cookie: <br>%v<br>
						<h1><a href="/expire">Delete the cookie</a></h1>
					`, c)
}
func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Println("No cookies can be read ! ")
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
