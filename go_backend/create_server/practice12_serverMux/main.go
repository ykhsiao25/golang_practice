package main

//Best Practice(use HandleFunc)
//HandleFunc() attaches to defaultServeMux
//use func() as parameters
import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}
func main() {

	//c, d are handlers
	//如果 想要 Execute /dog/something，這個'/' """一定"""要加!!!
	http.HandleFunc("/dog/", d) //沒差，加了'/'，在Url若沒加上'/'，會自動幫你添加且自動執行
	// http.HandleFunc("/cat", c)
	http.Handle("/cat", http.HandlerFunc(c)) //後面那不叫座call function，這叫coverting function type

	http.ListenAndServe(":8080", nil)
}
