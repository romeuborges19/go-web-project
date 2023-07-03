package controller

import (
	"cserver/domain"
	"net/http"
)

func (c *Controller) RegisterForm(w http.ResponseWriter, r *http.Request) {
	userInfo, logged := c.GetSessionData(r)

	tmpl.ExecuteTemplate(w, "form_register.html", struct {
		Logged bool
		User domain.Person
	} {
		Logged: logged,
		User: userInfo,
	})
}
