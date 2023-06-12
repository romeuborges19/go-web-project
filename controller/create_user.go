package controller

import (
	"cserver/domain"
	"log"
	"net/http"
)

func (c *Controller) CreateUser (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userInfo := domain.Person{
		Username: 	r.PostFormValue("username"),
		Email: 		r.PostFormValue("email"),
		Password: 	r.PostFormValue("password"),
	}

	_, err := c.userService.CreateUser(userInfo, c.db)	
	if err != nil {
		log.Fatal(err)
		return
	}
	http.Redirect(w, r, "/success", http.StatusSeeOther)
	return
}
