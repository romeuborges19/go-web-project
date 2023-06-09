package controller

import (
	"cserver/domain"
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	var questions []domain.Question
	questions, err := c.questionService.GetQuestions(c.db)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = r.Cookie("session-name")

	var logged bool
	if err != nil {
		logged = false
	} else {
		logged = true
	}

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal(err)
		return
	}

	username := fmt.Sprint(session.Values["username"])
	userInfo, err := c.userService.GetUserByUsername(username, c.db)


	tmpl.ExecuteTemplate(w, "index.html", struct {
		Logged bool
		Questions []domain.Question
			User domain.Person
		}{ 	
			Logged: logged,
			Questions: questions,
			User: userInfo,
	})
}
