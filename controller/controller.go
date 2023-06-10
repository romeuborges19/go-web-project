package controller

import (
	"cserver/service"
	"database/sql"
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
