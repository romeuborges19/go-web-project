package controller

import "net/http"

func (c *Controller) LoginForm(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "form_login.html", nil)
}
