package controller

import (
	"cserver/domain"
	"log"
	"net/http"
)

func (c *Controller) QuestionForm(w http.ResponseWriter, r *http.Request){
	userInfo, logged := c.GetSessionData(r)
	categories, err := c.categoryService.GetCategories(c.db)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "form_question.html", struct {
		Logged bool
		User domain.Person
		Categories []domain.Category
	} {
		Logged: logged,
		User: userInfo,
		Categories: categories,
	})
}
