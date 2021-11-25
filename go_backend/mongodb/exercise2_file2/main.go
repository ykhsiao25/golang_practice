package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/exercise2_file2/controllers"
	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/exercise2_file2/models"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[bson.ObjectId]models.User {
	m := map[bson.ObjectId]models.User{}
	return m
}
