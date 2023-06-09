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
	fmt.Println("newdb call 1")
	DB, err := sql.Open("postgres", dbconn)
	fmt.Println("newdb call 2")
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
