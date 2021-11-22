package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}
func bar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(404), http.StatusBadRequest)
	}
	fmt.Fprintln(res, "YOUR COOKIE: ", c)
}
