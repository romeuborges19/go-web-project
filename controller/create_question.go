package controller

import (
	"cserver/domain"
	"fmt"
	"log"
	"net/http"
	"time"
	// "github.com/gorilla/sessions"
)

func (c *Controller) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_, err := r.Cookie("session-name")
	if err != nil {
		http.Redirect(w, r, "/author/", http.StatusSeeOther)	
	} else {
		session, err := store.Get(r, "session-name")

		if err != nil {
			log.Fatal(err)
			return
		}

		username := fmt.Sprint(session.Values["username"])

		question := domain.Question{
			Title: r.PostFormValue("title"),
			Description: r.PostFormValue("description"),
			CreatedAt: time.Now(),
		}

		_, err = c.questionService.CreateQuestion(question, username, c.db)
	}
}
