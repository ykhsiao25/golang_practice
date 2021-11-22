package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println("NO Form File")
		}
		defer mf.Close()

		h := sha1.New()
		_, err = io.Copy(h, mf)
		if err != nil {
			log.Fatalln(err)
		}
		ext := strings.Split(fh.Filename, ".")[1]

		//fomat字串(按照 %<type>去組成字串)，並return string
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		c = appendValue(res, c, fname)

		// create the file
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		dirname := filepath.Join(pwd, "public", "pics")
		err = os.MkdirAll(dirname, os.ModeDir)
		if err != nil {
			log.Fatalln(err)
		}
		absFname := filepath.Join(dirname, fname)

		//這邊err可以直接定義(因為f還沒被定義)
		f, err := os.Create(absFname)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		//只有都不需要被定義的時候，才會用 "=" 而不是":="
		//一定要寫Seek() 不然可能檔案複製不完全
		_, err = mf.Seek(0, 0)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = io.Copy(f, mf)
		if err != nil {
			log.Fatalln(err)
		}
	}
	s := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(res, "index.gohtml", s)
}

func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		uid := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(res, c)
	}
	return c
}

func appendValue(res http.ResponseWriter, c *http.Cookie, fn string) *http.Cookie {
	// Values
	s := c.Value
	if !strings.Contains(s, fn) {
		s += "|" + fn
	}
	c.Value = s
	http.SetCookie(res, c)

	return c
}
