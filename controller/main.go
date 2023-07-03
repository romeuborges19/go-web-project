package controller

import (
	"cserver/domain"
	"log"
	"net/http"
)

func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	var questions []domain.Question
	questions, err := c.questionService.GetQuestions(c.db)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, logged := c.GetSessionData(r)

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
