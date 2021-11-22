package main

import (
	"fmt"
	"net/http"
)

func getUser(req *http.Request) user {
	var user1 user
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("No cookie here")
		return user1
	}
	userid, ok1 := dbSessions[c.Value]
	if !ok1 {
		fmt.Println("No uuid here")
		return user1
	}
	user1, ok2 := dbUsers[userid]
	if !ok2 {
		fmt.Println("No userinfo here")
		return user1
	}
	return user1
}

//use the cookie to judge whether logging in or not
func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	userid, ok1 := dbSessions[c.Value]
	if !ok1 {
		fmt.Println("No uuid here")
		return false
	}
	_, ok2 := dbUsers[userid]
	return ok2
}
