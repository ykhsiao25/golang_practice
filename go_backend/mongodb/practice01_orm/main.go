package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoesToEleven/golang-web-dev/042_mongodb/02_json/models"
	"github.com/julienschmidt/httprouter"
)

// client GET index url -> server回index(首頁)頁面 -> client點下超連結 -> GET 超連結url -> server回超連結response
func main() {
	r := httprouter.New()
	//我猜是有實作 ServeHttp() (這樣就可以做為Handler interface帶入ListenAndServe()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe(":8080", r)
}
func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
func getUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user1 := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"), //first matched  Param, or empty string
	}
	bs, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("Conent-Type", "application/json")
	// res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(bs))

}

//他並不是單純createUser，他是parse req之後改user attribute再return
func createUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user1 := models.User{}
	json.NewDecoder(req.Body).Decode(&user1)
	//這邊就是要確保create時 有確定被修改
	user1.Id = "007"
	bs, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("Content-Type", "application/json")
	// res.WriteHeader(http.StatusCreated) //201
	fmt.Fprintln(res, string(bs))
}

func deleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// res.WriteHeader(http.StatusOK) //200
	fmt.Fprintln(res, "this is deleteUser func!")
}
