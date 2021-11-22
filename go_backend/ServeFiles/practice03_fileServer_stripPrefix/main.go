package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	//這邊的resources/ 都是在browser上的url，他真正map到的 FileServer是 ./assets(所以local端和resources/這url本身無關)
	//http.StripPrefix()的用途是，前面Handle("<url>")依然是client要取得資源的url，但要取得資源必須map到local端路徑
	//而後面http.Dir("<local端路徑>")就是url map到的路徑 (總之StripPrefix()不是讓client不打/resources就能取得資源，他是幫後面FileServer()重定向)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//<img src="<url>" 是指browser的部分，和local端無關(他不是從./resources/toby.jpg 拿圖)
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*

./assets/toby.jpg

*/
