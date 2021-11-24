package main

import (
	"fmt"
	"net/http"
	"test1/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
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
