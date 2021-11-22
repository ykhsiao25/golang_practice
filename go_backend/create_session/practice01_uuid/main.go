package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	//寫法都是先取cookie，取不到再自己新增
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		//httpOnly: JavaScript 的Document.cookie (en-US) API 無法取得 HttpOnly cookies
		//抵禦攻擊者利用 Cross-Site Scripting (XSS) 手法來盜取用戶身份
		//secure:  有設定時，Cookie只能透過https的方式傳輸。
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true, //代表只可以透過Https取得
			Path:     "/",
		}
		http.SetCookie(res, c)
	}

	// fmt.Fprintln(res, c)
	fmt.Println(c)
}
