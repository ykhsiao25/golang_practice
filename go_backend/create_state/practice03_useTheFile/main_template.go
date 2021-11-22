package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

// 拿file -> 讀出file的[]byte ->  create file在server端(本機) -> 將[]byte寫入file -> 在server 印出string([]byte)
//因為s可為空字串，所以可以在method之外
func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("REQUEST PROCESS")
	var s string
	if req.Method == http.MethodPost {
		//formfile 的multipart.file 和 os.File是不同的(前者為interface，後者為type)
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		defer f.Close()
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		s = string(bs)

		//store file in server
		f2, err2 := os.Create(filepath.Join("./user/", h.Filename))
		if err2 != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		defer f2.Close()

		_, err3 := f2.Write(bs)
		if err3 != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.gohtml", s)
}
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
