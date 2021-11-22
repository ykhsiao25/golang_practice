package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

//用cookie 拿 uid(create session) -> dbSessions找username -> dbUsers用username拿user
//若dbSession沒有session，就找不到userid(這邊是利用cookie拿user)
func getUser(res http.ResponseWriter, req *http.Request) user {
	var user1 user
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

	if sess, ok := dbSessions[c.Value]; ok {
		sess.lastActivity = time.Now()
		dbSessions[c.Value] = sess
		user1 = dbUsers[sess.Username]
	}

	return user1
}

// 把all Sessions的 uid and username都print出來
func showSessions() {
	fmt.Println("*********")
	for uid, sess := range dbSessions {
		fmt.Println(uid, sess.Username)
	}
	fmt.Println("")
}

// alreadyLoggedInt == cookie存在 + session存在 + user存在
func alreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	//cookie不存在
	c, err1 := req.Cookie("session")
	if err1 != nil {
		return false
	}
	// session不存在
	sess, ok := dbSessions[c.Value]
	if !ok {
		return false
	} else {
		dbSessions[c.Value] = sess
		http.SetCookie(res, c)
	}
	_, ok = dbUsers[sess.Username]
	return ok
}

// delete all > dbSessionLength sessions + renew dbSessionsCleaned to time.Now()
func cleanSessions() {
	fmt.Println("Before Clean the session")
	for uid, sess := range dbSessions {
		if time.Since(sess.lastActivity) > (time.Second * 30) {
			fmt.Println("session time out 2")
			delete(dbSessions, uid)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}
