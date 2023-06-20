package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options = &sessions.Options{
		Path: "/",
		MaxAge: 86400 * 7,
		HttpOnly: true,
	}


	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	correctPassword, err := c.userService.GetPasswordByUsername(username, c.db) 
	if err != nil {
		http.Redirect(w, r, "/failure", http.StatusSeeOther)
		log.Fatal(err)
	}
	if username != "" {
		session.Values["username"] = username
	}
	err = sessions.Save(r, w)

	if password == correctPassword {
		http.Redirect(w, r, "/", http.StatusSeeOther)	
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
