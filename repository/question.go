package repository

import (
	"cserver/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type QuestionQuery interface {
	GetQuestions (db *sql.DB) ([]domain.Question, error)
	CreateQuestion (question domain.Question, db *sql.DB) (int, error)
	GetQuestionByID(questionID int, db *sql.DB) (domain.Question, error)
}

type questionQuery struct {}

func (q *questionQuery) CreateQuestion (question domain.Question, db *sql.DB) (int, error){
	query := `INSERT INTO "question"("title", "description", "id_autor") VALUES ($1, $2, $3)`

	_, err := db.Exec(query, question.Title, question.Description, question.AuthorID)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return 1, nil
}

func (q *questionQuery)	GetQuestions (db *sql.DB) ([]domain.Question, error) {
	fmt.Println("getquestions call")
	query := `SELECT * FROM "question"`
	fmt.Println("getquestions call 1")
	rows, err := db.Query(query)
	fmt.Println("getquestions call 2")
	if err != nil {
	fmt.Println("getquestions call error")
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	var questions []domain.Question
	var question domain.Question
	for rows.Next() {
		err := rows.Scan(&question.ID, &question.Title, &question.Description, &question.AuthorID)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		questions = append(questions, question)
	}
	
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return questions, nil
}

func (q *questionQuery) GetQuestionByID(questionID int, db *sql.DB) (domain.Question, error){
	query := `SELECT * FROM "question" WHERE id = $1`

	var id, authorID int
	var title, description string

	err := db.QueryRow(query, questionID).Scan(&id, &title, &description, &authorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Question{}, errors.New("user not found")
		}
		return domain.Question{}, err
	}

	questionInfo := domain.Question{
		ID: id,
		Title: title,
		Description: description,
		AuthorID: authorID,
	}
	return questionInfo, nil
}