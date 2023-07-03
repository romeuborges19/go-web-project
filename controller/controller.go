package controller

import (
	"cserver/service"
	"database/sql"
	"html/template"
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
