package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	res.Write([]byte(s))
}

func mshl(res http.ResponseWriter, req *http.Request) {
	p1 := Person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	res.Header().Set("Content-Type", "application/json")
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Write(bs)
}

func encd(res http.ResponseWriter, req *http.Request) {
	p1 := Person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(p1)
	if err != nil {
		log.Fatalln(err)
	}
}
