package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	//法一
	// count, err2 := strconv.Atoi(cookie.Value)
	// if err2 != nil {
	// 	log.Fatalln(err2)
	// }
	// count++
	// cookie.Value = strconv.Itoa(count)

	//法二
	count, err2 := strconv.ParseInt(cookie.Value, 10, 0)
	if err2 != nil {
		log.Fatalln(err2)
	}
	count++
	cookie.Value = strconv.FormatInt(int64(count), 10)

	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
