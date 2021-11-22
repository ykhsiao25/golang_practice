package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.html", "index2.html")
	if err != nil {
		log.Fatal("Error Parsing files", err)
	}
	// execute the specific template
	err = tpl.ExecuteTemplate(os.Stdout, "index2.html", nil)
	if err != nil {
		log.Fatal("Error Executing Template files", err)
	}
}
