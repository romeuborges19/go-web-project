package controller

import (
	"cserver/domain"
	"fmt"
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
	fmt.Println("erro1")
	questionInfo, err := c.questionService.GetQuestion(questionID, c.db)
	fmt.Println("erro2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("erro3")
	authorInfo, err := c.userService.GetUserByID(questionInfo.AuthorID, c.db)
	fmt.Println("erro4")
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
