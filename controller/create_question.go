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
		}

		username := fmt.Sprint(session.Values["username"])

		var current_timestamp time.Time 
		current_timestamp = time.Now()
		categoryName := r.PostFormValue("categories")
		
		category, err := c.categoryService.GetCategoryByName(categoryName, c.db)

		fmt.Println(category)
		if err != nil {
			log.Fatal(err)
		}

		question := domain.Question{
			Title: r.PostFormValue("title"),
			Description: r.PostFormValue("description"),
			Category: category,
			CreatedAt: &current_timestamp,
		}

		_, err = c.questionService.CreateQuestion(question, username, c.db)
		http.Redirect(w, r, "/", http.StatusSeeOther)	
	}
}
