package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ykhsiao25/golang_practice/go_backend/mongodb/exercise1/models"

	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session map[bson.ObjectId]models.User
}

func NewUserController(dbUsers map[bson.ObjectId]models.User) *UserController {
	return &UserController{dbUsers}
}

func (uc UserController) GetUser(res http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	// oid := bson.ObjectIdHex(id)

	// composite literal
	// user1 := models.User{}

	// Fetch user
	user1, ok := uc.session[bson.ObjectIdHex(id)]
	if !ok {
		res.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	bs, _ := json.Marshal(user1)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", bs)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user1 := models.User{}

	json.NewDecoder(r.Body).Decode(&user1)

	// create bson ID
	user1.Id = bson.NewObjectId()

	// store the user in mongodb
	// uc.session.DB("go-web-dev-db").C("users").Insert(u)
	uc.session[user1.Id] = user1

	bs, _ := json.Marshal(user1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", bs)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	// Delete user
	_, ok := uc.session[oid]
	if !ok {
		w.WriteHeader(404)
		return
	}
	delete(uc.session, oid)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
