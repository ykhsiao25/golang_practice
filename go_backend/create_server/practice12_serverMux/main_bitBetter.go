package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	//c, d are handlers
	//如果 想要 Execute /dog/something，這個'/' """一定"""要加!!!
	mux.Handle("/dog/", d) //沒差，加了'/'，在Url若沒加上'/'，會自動幫你添加且自動執行
	mux.Handle("/cat", c)

	//會自動註冊(defaultServeMux)
	http.ListenAndServe(":8080", mux)
}
