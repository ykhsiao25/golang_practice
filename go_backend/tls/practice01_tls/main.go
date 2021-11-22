package main

import "net/http"

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("This is an example server.\n"))
}
