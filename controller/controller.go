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
	userService 	service.UserService
	questionService service.QuestionService
	categoryService service.CategoryService
	db *sql.DB
	fs http.Handler
}

func NewController(userService service.UserService,	questionService service.QuestionService, categoryService service.CategoryService, db *sql.DB) *Controller {
	return &Controller{
		userService: userService,
		questionService: questionService,
		categoryService: categoryService,
		db: db,
		fs: http.FileServer(http.Dir("web/static")),
	}
}

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

func (c *Controller) LoginForm(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "form_login.html", nil)
}

func (c *Controller) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal("failed to get session", err)
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Fatal("failed to delete session", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
