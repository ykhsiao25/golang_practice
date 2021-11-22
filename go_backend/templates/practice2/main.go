package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	f, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("error parsing templates")
	}
	err = f.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
