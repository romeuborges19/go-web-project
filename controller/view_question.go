package controller

import (
	"cserver/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) QuestionView(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ID := vars["id"]

	questionID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	questionInfo, err := c.questionService.GetQuestion(questionID, c.db)
	if err != nil {
		log.Fatal(err)
	}
	authorInfo, err := c.userService.GetUserByID(questionInfo.AuthorID, c.db)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, logged := c.GetSessionData(r)

	tmpl.ExecuteTemplate(w, "view_question.html", struct {
		Logged 	 bool
		Question domain.Question
		User 	 domain.Person
		Author 	 domain.Person
	} {
		Logged: 	logged,
		Question: 	questionInfo,
		User: 		userInfo,
		Author: 	authorInfo,
	})
}
