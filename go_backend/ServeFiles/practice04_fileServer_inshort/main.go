package main

import (
	"log"
	"net/http"
)

func main() {
	//index.html 是special case 他會直接show畫面，其他就一般情況，show哪些files的超連結，並秀出contents
	log.Fatalln(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
