package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// directly get Users[userid]
func getUser(res http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		uid := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(res, c)

	// if the user exists already, get user
	var user1 user
	if sess, ok := dbSessions[c.Value]; ok {
		sess.lastActivity = time.Now()
		dbSessions[c.Value] = sess
		user1 = dbUsers[sess.userid]
	}
	return user1
}

//if Users[userid] no value, means no loggedIn
func alreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	sess, ok := dbSessions[c.Value]
	if ok {
		sess.lastActivity = time.Now()
		dbSessions[c.Value] = sess
	}
	_, ok = dbUsers[sess.userid]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(res, c)
	return ok
}
func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.userid)
	}
	fmt.Println("")
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, _ := range dbSessions {
		delete(dbSessions, k)
	}
	// dbSessionsCleaned = time.Now()
	fmt.Println("After CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}
