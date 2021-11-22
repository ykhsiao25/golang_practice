package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 拿file -> 讀出file的[]byte -> print出file
//因為s可為空字串，所以可以在method之外
func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("REQUEST PROCESS")
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		//因為input type 可以為 "file" 透過name取得這個file
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		defer f.Close()
		fmt.Println("\nFile: ", f, "\nHeader: ", h, "\nerr", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		s = string(bs)
	}
	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(res, `
	<form method="Post" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
