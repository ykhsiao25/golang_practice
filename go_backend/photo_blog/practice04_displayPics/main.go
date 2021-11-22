package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	//注意!! http.StripPrefix("/path", ...)前面的 url "務必一定"要寫成 "/path/" (後面一定要多一個/，不然抓不出檔案)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		check(err)
		defer mf.Close()

		h := sha1.New()
		_, err = io.Copy(h, mf)
		check(err)
		ext := strings.Split(fh.Filename, ".")[1]
		f_name := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		fmt.Println(f_name)
		c = appendValue(res, c, f_name)

		//Create the file
		pwd, err := os.Getwd()
		check(err)

		dirname := filepath.Join(pwd, "public", "pics")
		err = os.MkdirAll(dirname, os.ModeDir)
		check(err)

		f, err := os.Create(filepath.Join(dirname, f_name))
		check(err)
		defer f.Close()

		_, err = mf.Seek(0, 0)
		check(err)
		_, err = io.Copy(f, mf)
		check(err)
	}
	pics := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(res, "index.gohtml", pics[1:])
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

func appendValue(res http.ResponseWriter, c *http.Cookie, f_name string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, f_name) {
		s += "|" + f_name
	}
	c.Value = s
	http.SetCookie(res, c)
	return c
}
func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
