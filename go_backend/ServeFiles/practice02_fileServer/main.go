package main

import (
	"io"
	"net/http"
)

func main() {
	//在 路徑 "." 的所有files都作為FileServer
	//(他會給你 <a href=""> <filename> </a> 的連結，然後點下去會跳出file內容，包含code)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}
func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html; charset = UTF-8")
	io.WriteString(res, `<img src="/toby.jpg">`)
}
