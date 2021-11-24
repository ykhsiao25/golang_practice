package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/practice03_mongodbCRUD/models"
	// "github.com/GoesToEleven/golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(sess *mgo.Session) *UserController {
	return &UserController{session: sess}
}

//get id 轉成 oid -> 透過oid去db 找user -> response user
func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("Id")
	//judge 是否為oid
	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)

	user1 := models.User{}
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&user1); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
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

	// user1.Id = "007"
	user1.Id = bson.NewObjectId()
	// C == collection == RDBMS.Table
	uc.session.DB("go-web-dev-db").C("users").Insert(user1)

	bs, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) //201
	fmt.Fprintln(res, string(bs))
}
func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("Id")
	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "Delete User", oid)

}
