package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//這就單純取得 FormValue()的API而已(就html 那個form)
	// http://127.0.0.1:8080/?q=12 這樣打
	v := req.FormValue("q")
	v2 := req.FormValue("qq")
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(res, `
	<form method="POST">
	<input type="text" name="q">
	<input type="text" name="qq">
	<input type="submit">
	</form>
	<br>`+v+v2)
}
