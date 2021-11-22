package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err1 error

func main() {
	db, err1 = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test01?charset=utf8")
	check(err1)
	defer db.Close()

	err1 = db.Ping()
	check(err1)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err1 = http.ListenAndServe(":8080", nil)
	check(err1)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err1 = io.WriteString(res, "Sucessfully complete!")
	check(err1)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
