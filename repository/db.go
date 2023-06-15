package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type DAO interface{
	NewUserQuery() UserQuery
	NewQuestionQuery() QuestionQuery
	NewCategoryQuery() CategoryQuery
}

type dao struct{}

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbconn := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable user=%s password=%s", host, port, dbname, user, password)
	DB, err := sql.Open("postgres", dbconn)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return DB, err
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewQuestionQuery() QuestionQuery {
	return &questionQuery{}
}

func (d *dao) NewCategoryQuery() CategoryQuery {
	return &categoryQuery{}
}
