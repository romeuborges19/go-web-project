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
	fmt.Println("questionID: ", questionID)

	questionInfo, err := c.questionService.GetQuestion(questionID, c.db)
	if err != nil {
		log.Fatal(err)
		return
	}
	userInfo, logged := c.GetSessionData(r)

	tmpl.ExecuteTemplate(w, "view_question.html", struct {
		Logged bool
		User domain.Person
		Question domain.Question
	} {
		Logged: logged,
		User: userInfo,
		Question: questionInfo,
	})
}
