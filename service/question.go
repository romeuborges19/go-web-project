package service

import (
	"cserver/domain"
	"cserver/repository"
	"database/sql"
	// "fmt"
	"log"
)

type QuestionService interface {
	GetQuestions(db *sql.DB) ([]domain.Question, error)
	CreateQuestion (question domain.Question, authorUsername string, db *sql.DB) (int, error)
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
	var questions []domain.Question
	var err error
	questions, err = q.dao.NewQuestionQuery().GetQuestions(db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return questions, nil
}

func (q *questionService) GetQuestion(questionID int, db *sql.DB) (domain.Question, error){
	questionInfo, err := q.dao.NewQuestionQuery().GetQuestionByID(questionID, db)
	if err != nil {
		log.Fatal()
		return domain.Question{}, err
	}
	return questionInfo, nil
}
