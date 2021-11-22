package main

import "net/http"

func getUser(req *http.Request) user {
	var user1 user
	c, err := req.Cookie("session")
	if err != nil {
		return user1
	}
	userid := dbSessions[c.Value]
	user1 = dbUsers[userid]
	return user1
}
func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	userid, ok1 := dbSessions[c.Value]
	if !ok1 {
		return false
	}
	_, ok2 := dbUsers[userid]
	return ok2
}
