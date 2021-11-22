package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//這就單純取得 FormValue()的API而已(就html 那個form)
	// http://127.0.0.1:8080/?q=12 這樣打
	v := req.FormValue("q")
	fmt.Fprintln(res, "Do my search: "+v)

}
