package main

import (
	"fmt"
	"net/http"
)

func getUser(req *http.Request) user {
	var user1 user
	c, err1 := req.Cookie("session")
	if err1 != nil {
		fmt.Println("NO cookie yet! in  getUser()")
		return user1
	}
	if uid, ok := dbSessions[c.Value]; ok {
		return dbUsers[uid]
	}
	return user1
}
func alreadyLoggedIn(req *http.Request) bool {
	c, err1 := req.Cookie("session")
	if err1 != nil {
		fmt.Println("NO cookie yet! in  alreadyLoggedIn()")
		return false
	}
	userid, _ := dbSessions[c.Value]
	_, ok := dbUsers[userid]
	return ok
}
