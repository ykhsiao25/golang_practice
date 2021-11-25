package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/practice03_mongodbCRUD/controllers"

	// "github.com/GoesToEleven/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"

	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())
	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	sess, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatalln(err)
	}
	return sess
}
