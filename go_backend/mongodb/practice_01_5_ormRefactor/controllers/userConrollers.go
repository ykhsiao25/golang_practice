package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user1 := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	bs, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", bs)
}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user1 := models.User{}

	json.NewDecoder(req.Body).Decode(&user1)
	//這樣也可以
	// bs, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// json.Unmarshal(bs, &user1)

	user1.Id = "007"

	bs, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(res, "%s\n", bs)
}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(res, "Write code to delete user\n")
}
