package service

import (
	"cserver/domain"
	"cserver/repository"
	"database/sql"
	"fmt"

	// "fmt"
	"log"
)

type QuestionService interface {
	CreateQuestion (question domain.Question, authorUsername string, db *sql.DB) (int, error)
	GetQuestions(db *sql.DB) ([]domain.Question, error)
	GetQuestion(questionID int, db *sql.DB) (domain.Question, error)
}

type questionService struct {
	dao repository.DAO
}

func NewQuestionService(dao repository.DAO) QuestionService {
	return &questionService{dao: dao}
} 

func (q *questionService) CreateQuestion (question domain.Question, authorUsername string, db *sql.DB) (int, error) {

	userID, err := q.dao.NewUserQuery().GetIDByUsername(authorUsername, db)
	question.AuthorID = userID
	_, err = q.dao.NewQuestionQuery().CreateQuestion(question, db)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return 1, nil
}

func (q *questionService) GetQuestions(db *sql.DB) ([]domain.Question, error){
	fmt.Println("erro service 0")
	var questions []domain.Question
	var err error
	fmt.Println("erro service 1")
	questions, err = q.dao.NewQuestionQuery().GetQuestions(db)
	fmt.Println("erro service 2")
	if err != nil {
		log.Fatal(err)
	}
	return questions, nil
}

func (q *questionService) GetQuestion(questionID int, db *sql.DB) (domain.Question, error){
	questionInfo, err := q.dao.NewQuestionQuery().GetQuestionByID(questionID, db)
	if err != nil {
		log.Fatal(err)
	}
	return questionInfo, nil
}
