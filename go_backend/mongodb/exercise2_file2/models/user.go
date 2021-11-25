package models

import (
	"encoding/json"
	"net/http"
	"os"

	// "github.com/ykhsiao25/golang_practice/go_backend/mongodb/exercise2_file2/controllers"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

// Id was of type string before
//參數不可以用uc，否則import cycle not allowed
func UpdateUser(res http.ResponseWriter, dbUsers *map[bson.ObjectId]User) {
	f, err := os.Create("dbUsers/data.json")
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	json.NewEncoder(f).Encode(*dbUsers)
}

func LoadUser(res http.ResponseWriter, dbUsers *map[bson.ObjectId]User) {
	f, err := os.Open("dbUsers/data.json")
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	json.NewDecoder(f).Decode(dbUsers)
}
