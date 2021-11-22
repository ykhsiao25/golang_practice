package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}
func foo(res http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)

	//一定要特別注意，Read() retrun 的 error可能是 EOF，而不是nil，但這代表要成功讀完，一定要特別注意!!
	_, err := req.Body.Read(bs)
	if err.Error() != "EOF" {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
	//或乾脆不要取error
	// req.Body.Read(bs)

	body := string(bs)
	err2 := tpl.ExecuteTemplate(res, "index.html", body)
	if err2 != nil {
		http.Error(res, err2.Error(), http.StatusInternalServerError)
		log.Fatalln(err2)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
