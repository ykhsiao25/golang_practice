package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// index.gohtml create
	// name := "Peter H"
	// tpl := `
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	// <meta charset="UTF-8">
	// <title>Hello World</title>
	// </head>
	// <body>
	// <h1> ` + name + `</h1>
	// </body>
	// </html>
	// `
	// fmt.Println(tpl)

	tpl, err := template.ParseFiles("index.gohtml", "index.html")
	if err != nil {
		log.Fatal("Error parsing fles", err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		// fmt.Println("Error creating file ", err)
		// os.Exit(1)
		log.Fatal("Error creating file", err)
	}
	defer f.Close()

	//copy strings to a file
	// io.Copy(f, strings.NewReader(tpl))
	err = tpl.Execute(f, nil)
	if err != nil {
		log.Fatal(err)
	}
}
