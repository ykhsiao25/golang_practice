package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//只要implement ServeHTTP(http.ResponseWriter, *http.Request) 就可以作為 interface handler{} 作為params代入
func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Any code you want in this func")
}

//機制是這樣的
// 實現handler method() -> 可用handler type的func() (ListenAndServe) -> 把handler and ip addr 包成 Server type ->
// listen() -> (listener) accept() -> (conn) serve(這個就執行)
func main() {
	var d hotdog
	//:8080 應該是ip:port, 因為預設localhost, 所以不寫
	http.ListenAndServe(":8080", d)
}
