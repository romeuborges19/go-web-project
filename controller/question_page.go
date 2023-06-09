package controller

import (
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

	tmpl.ExecuteTemplate(w, "question_page.html", questionInfo)
}
