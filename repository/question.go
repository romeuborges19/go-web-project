package repository

import (
	"cserver/domain"
	"database/sql"
	"errors"
	"log"
)

type QuestionQuery interface {
	GetQuestions (db *sql.DB) ([]domain.Question, error)
	CreateQuestion (question domain.Question, db *sql.DB) (int, error)
	GetQuestionByID(questionID int, db *sql.DB) (domain.Question, error)
}

type questionQuery struct {}

func (q *questionQuery) CreateQuestion (question domain.Question, db *sql.DB) (int, error){
	query := `INSERT INTO "question"("title", "description", "author_id", "category_id") VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, question.Title, question.Description, question.Author.ID, question.Category.ID)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return 1, nil
}

func (q *questionQuery)	GetQuestions (db *sql.DB) ([]domain.Question, error) {
	query := `SELECT * FROM "question" ORDER BY "id" DESC`
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var questions []domain.Question
	var question domain.Question

	for rows.Next() {
		err := rows.Scan(
			&question.ID, 
			&question.Title, 
			&question.Description, 
			&question.Author.ID, 
			&question.Category.ID,
			&question.CreatedAt, 
			&question.ModifiedAt, 
		)

		if err != nil {
			log.Fatal(err)
		}

		query = `SELECT "first_name", "last_name", "username" FROM "person" WHERE id = $1`
		err = db.QueryRow(query, question.Author.ID).Scan(
			&question.Author.FirstName,
			&question.Author.LastName,
			&question.Author.Username,
		)
		if err != nil {
			log.Fatal(err)
		}

		query = `SELECT "name" FROM "category" WHERE id = $1`
		err = db.QueryRow(query, question.Category.ID).Scan(
			&question.Category.Name,
		)
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, question)
	}
	
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return questions, nil
}

func (q *questionQuery) GetQuestionByID(questionID int, db *sql.DB) (domain.Question, error){
	query := `SELECT * FROM "question" WHERE id = $1`

	var question domain.Question

	err := db.QueryRow(query, questionID).Scan(
			&question.ID, 
			&question.Title, 
			&question.Description, 
			&question.Author.ID, 
			&question.Category.ID,
			&question.CreatedAt, 
			&question.ModifiedAt, 
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Question{}, errors.New("user not found")
		}
		return domain.Question{}, err
	}

	return question, nil
}
