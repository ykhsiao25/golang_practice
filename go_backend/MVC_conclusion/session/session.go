package session

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/ykhsiao25/golang_practice/go_backend/MVC_conclusion/models"
)

const Length int = 30

//一定要寫在這裡，不然和controllers會import cycle
var DbSessions = map[string]models.Session{}
var DbUsers = map[string]models.User{}
var LastCleaned = time.Now()

func GetUser(res http.ResponseWriter, req *http.Request) models.User {
	//create the session
	var user1 models.User
	ck, err := req.Cookie("session")
	if err != nil {
		uid := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
	}
	ck.MaxAge = Length
	http.SetCookie(res, ck)

	//update the session
	sess := DbSessions[ck.Value]
	sess.LastActivity = time.Now()
	DbSessions[ck.Value] = sess
	// get the user
	user1 = DbUsers[sess.UserName]

	return user1
}

func AlreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	//no cookie
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//update the cookie
	ck.MaxAge = Length
	http.SetCookie(res, ck)

	//no session
	sess, ok := DbSessions[ck.Value]

	// update the session
	sess.LastActivity = time.Now()
	DbSessions[ck.Value] = sess
	if !ok {
		return false
	}

	//no user
	if _, ok = DbUsers[sess.UserName]; !ok {
		return false
	}

	return true
}
func Show() {
	fmt.Println("********")
	for uid, sess := range DbSessions {
		fmt.Println("uid: ", uid, " and sess: ", sess)
	}
	fmt.Println("********")
}

func Clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for oid, sess := range DbSessions {
		if time.Since(sess.LastActivity) > (time.Second * 30) {
			delete(DbSessions, oid)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}
