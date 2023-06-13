package main

import (
	"cserver/controller"
	"cserver/repository"
	"cserver/service"
	"fmt"
	"net/http"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	// DATABASE
	db, err := repository.NewDB()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()
	mux := mux.NewRouter()

	// Initializing all services
	dao := repository.NewDAO()
	userService := service.NewUserService(dao)
	questionService := service.NewQuestionService(dao)
	c := controller.NewController(userService, questionService, db)

	// Parsing static files to the server
	fs := http.FileServer(http.Dir("web/static"))
	mux.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static", fs))

	// Handling the website routes
	mux.Handle("/author/web/static/", http.StripPrefix("/author/web/static/", fs))
	mux.HandleFunc("/", c.Handler)
	mux.HandleFunc("/author/register", c.RegisterForm)
	mux.HandleFunc("/author/register-process", c.CreateUser)
	mux.HandleFunc("/author/", c.LoginForm)
	mux.HandleFunc("/author/login", c.Login)
	mux.HandleFunc("/author/ask", c.QuestionForm)
	mux.HandleFunc("/author/question", c.CreateQuestion)
	mux.HandleFunc("/delete-session", c.DeleteSession)
	mux.HandleFunc("/question/{id}", c.QuestionView)
	
	http.ListenAndServe(":8080", context.ClearHandler(mux))
}



