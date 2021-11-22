package main

import (
	"html/template"
	"io"
	"net/http"
)

func main() {
	//只有 '/'是給Web url用的，一般local要用 "./"
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogpic) // 這個就是直接給你放圖用的(可以直接serveFiles()
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}
func dog(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./index.html"))
	tpl.ExecuteTemplate(res, "index.html", nil)
}
func dogpic(res http.ResponseWriter, req *http.Request) {
	//這裡不能寫成 /dog.jpg 要馬寫成 "./dog.jpg" or "dog.jpg"
	http.ServeFile(res, req, "./dog.jpg")
}
