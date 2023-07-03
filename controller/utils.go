package controller

import (
	"cserver/domain"
	"fmt"
	"log"
	"net/http"
)


func (c *Controller) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal("failed to get session", err)
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Fatal("failed to delete session", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CheckLogin(r *http.Request) bool {
	_, err := r.Cookie("session-name")

	var logged bool
	if err != nil {
		logged = false
	} else {
		logged = true
	}
	return logged
}

func (c *Controller) GetSessionData(r *http.Request) (domain.Person, bool) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal(err)
	}

	var username string
		var userInfo domain.Person

	logged := CheckLogin(r)

	if logged {
		username = fmt.Sprint(session.Values["username"])
		userInfo, err = c.userService.GetUserByUsername(username, c.db)

		if err != nil {
			log.Fatal(err)
		}
	}
	return userInfo, logged
}
