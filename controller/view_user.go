package controller

import (
	"cserver/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) UserView(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ID := vars["id"]

	userID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := c.userService.GetUserByID(userID, c.db)
	_, logged := c.GetSessionData(r)

	tmpl.ExecuteTemplate(w, "view_user.html", struct {
		Logged bool
		User domain.Person
	}{
		Logged: logged,
		User: userInfo,
	})
}
