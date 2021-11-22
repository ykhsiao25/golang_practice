package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html; charset = UTF-8")
	io.WriteString(res, `
	<img src="/toby.jpg">
	`)
}

// func dogpic(res http.ResponseWriter, req *http.Request) {
// 	f, err := os.Open("toby.jpg")
// 	if err != nil {
// 		http.Error(res, "File not found", 404)
// 		return
// 	}
// 	defer f.Close()

// 	//ServeFile 法一
// 	// io.Copy(res, f)

// 	//ServeFile 法二
// 	fi, err2 := f.Stat()
// 	if err2 != nil {
// 		http.Error(res, "file not found", 404) //http.StatusNotFound
// 		return
// 	}

// 	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
// }

//ServeFile 法三 (個人覺得最佳)
func dogpic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}
