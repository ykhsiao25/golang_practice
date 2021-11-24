package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/practice02_mongodb/models"
	// "github.com/GoesToEleven/golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(sess *mgo.Session) *UserController {
	return &UserController{session: sess}
}
func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user1 := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}
	bs, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(bs))
}
func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user1 := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user1)
	if err != nil {
		log.Fatalln(err)
	}
	user1.Id = "007"
	bs, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) //201
	fmt.Fprintln(res, string(bs))
}
func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	res.WriteHeader(http.StatusOK)
	fmt.Println("this is the DeleteUser()")
}
