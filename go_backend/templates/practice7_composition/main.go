package main

import (
	"log"
	"os"
	"text/template"
)

//兩者相同意思
// type course struct {
// 	Number, Name, Units string
// }
type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Season  string
	Courses []course
}

// 注意這邊是叫做composition，不是繼承，所以不可以直接使用semester的property (ex: Season)
type year struct {
	Spring, Summer, Fall semester
}

var tpl1 *template.Template

func init() {
	//注意是ParseFiles() or ParseGlob()
	tpl1 = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	y := year{
		Fall: semester{
			Season: "Fall",
			Courses: []course{
				{"CSCI-40", "Introduction to Programming in Go", "4"},
				{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Season: "Spring",
			Courses: []course{
				// redundant type from array, slice, or map composite literal
				//因為已經宣告過是course，所以不要寫 course{}，直接{}就好
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming with Go", "5"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}
	err1 := tpl1.ExecuteTemplate(os.Stdout, "index.html", y)
	if err1 != nil {
		log.Fatal(err1)
	}
}
