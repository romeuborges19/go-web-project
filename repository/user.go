package repository

import (
	"cserver/domain"
	"database/sql"
	"errors"

	"log"
)

type UserQuery interface {
	CreateUser (person domain.Person, db *sql.DB) (int64, error)
	GetPasswordByUsername (username string, db *sql.DB) (string, error)
	GetIDByUsername (username string, db *sql.DB) (int, error)
	GetUserByID (userID int, db *sql.DB) (domain.Person, error)
}

type userQuery struct {}

func (u *userQuery) CreateUser (user domain.Person, db *sql.DB) (int64, error) {
	query := `INSERT INTO "person"("first_name", "last_name", "username", "email", "password") VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, 
	user.FirstName, 
	user.LastName,
	user.Username,
	user.Email, 
	user.Password)

	if err != nil {
		log.Fatal(err)
	}
	return 1, nil

}

func (u *userQuery) GetUserByID (userID int, db *sql.DB) (domain.Person, error){
	query := `SELECT * FROM "person" WHERE "id" = $1`


	var user domain.Person

	err := db.QueryRow(query, userID).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	if err == sql.ErrNoRows {
		return domain.Person{}, errors.New("user not found")
	}

	return user, nil
}

func (u *userQuery) GetPasswordByUsername (username string, db *sql.DB) (string, error) {
	query := `SELECT "password" FROM "person" WHERE "username" = $1`

	var password string
	row := db.QueryRow(query, username).Scan(&password)

	if row == sql.ErrNoRows {
		return password, errors.New("user not found")
	}

	return password, nil
}

func (u *userQuery) GetIDByUsername (username string, db *sql.DB) (int, error) {
	query := `SELECT "id" FROM "person" WHERE "username" = $1`

	var userID int
	row := db.QueryRow(query, username).Scan(&userID)

	if row == sql.ErrNoRows {
		return 0, errors.New("user not found")
	}
	return userID, nil
}
