package controller

import (
	"cserver/domain"
	"cserver/service"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var tmpl *template.Template
var store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	tmpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

type Controller struct{
	userService service.UserService
	questionService service.QuestionService
	db *sql.DB
	fs http.Handler
}

func NewController(userService service.UserService,	questionService service.QuestionService, db *sql.DB) *Controller {
	return &Controller{
		userService: userService,
		questionService: questionService,
		db: db,
		fs: http.FileServer(http.Dir("web/static")),
	}
}

func (c *Controller) RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func (c *Controller) QuestionPage(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "question_form.html", nil)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func (c *Controller) TemplateTest(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "processor.html", nil)
}

func (c *Controller) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal("failed to get session", err)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Fatal("failed to delete session", err)
		return 
	}

}

func CheckLogin(r *http.Request) bool {
	_, err := r.Cookie("session-name")

	var logged bool
	if err != nil {
		logged = false
	} else {
		logged = true
	}
	return logged
}

func (c *Controller) GetSessionData(r *http.Request) (domain.Person, bool) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal(err)
	}

	var username string
	var userInfo domain.Person

	logged := CheckLogin(r)

	if logged {
		username = fmt.Sprint(session.Values["username"])
		userInfo, err = c.userService.GetUserByUsername(username, c.db)

		if err != nil {
			log.Fatal(err)
		}
	}
	return userInfo, logged
}
