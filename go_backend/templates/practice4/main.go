package main

import (
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

//package-level var
var tpl *template.Template

func createHtml() {
	// tpl
	name := "Peter H"
	tplStr := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1> ` + name + `</h1>
	</body>
	</html>
	`
	f, err := os.Create("index1.html")
	if err != nil {
		// fmt.Println("Error creating file ", err)
		// os.Exit(1)
		log.Fatal("Error creating file", err)
	}
	io.Copy(f, strings.NewReader(tplStr))
	defer f.Close()
}

//import package就會執行
func init() {
	createHtml()
	copyHtml("index1.html", "index2.html")
	// after glob return template, err, if err != nil, then panic
	// with template os, we don't have to do error checking
	tpl = template.Must(template.ParseGlob("./*.html"))
	// tpl, err := template.ParseFiles("index.html", "index2.html")
}

func copyHtml(fileName string, fileName2 string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	f2, err2 := os.Open(fileName2)
	if err2 != nil {
		log.Fatal(err2)
	}
	io.Copy(f, f2)
	defer f.Close()
	defer f2.Close()
}

func main() {
	//不要 用 "./index.html" 要用 "index.html"
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatal("error Executing files ", err)
	}
	err2 := tpl.ExecuteTemplate(os.Stdout, "index2.html", `Release self-focus; embrace other-focus.`)
	if err2 != nil {
		log.Fatal(err2)
	}
}
