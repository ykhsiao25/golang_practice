package main

import (
	"fmt"
	"net/http"

	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/practice02_mongodb/controllers"
	// "github.com/GoesToEleven/golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	//注意，這邊一定不可以加上 "/"，不然會找不到
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("127.0.0.1:8080", r)
}

func index(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	s := `<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>INDEX</title>
			</head>
			<body>
				<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
			</body>
		</html>
		`
	res.Header().Set("Content-Type", "text/html; charset=UFT-8")
	//這應該可以不用寫
	// res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, s) //範例用 res.Write([]byte(s))
}
func getSession() *mgo.Session {
	sess, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return sess
}
