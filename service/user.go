package service

import (
	"cserver/domain"
	"cserver/repository"
	"database/sql"

	"log"
)

type UserService interface{
	CreateUser (person domain.Person, db *sql.DB) (int64, error)
	GetPasswordByUsername (username string, db *sql.DB) (string, error)
	GetUserByUsername (username string, db *sql.DB) (domain.Person, error)
}

type userService struct{
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao,}
}

func (u *userService) CreateUser (person domain.Person, db *sql.DB) (int64, error) {
	personID, err := u.dao.NewUserQuery().CreateUser(person, db)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return personID, nil
}

func (u *userService) GetPasswordByUsername (username string, db *sql.DB) (string, error){
	password, err := u.dao.NewUserQuery().GetPasswordByUsername(username, db)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return password, nil
}

func (u *userService) GetUserByUsername (username string, db *sql.DB) (domain.Person, error) {
	userID, err := u.dao.NewUserQuery().GetIDByUsername(username, db)
	if err != nil {
		log.Fatal(err)
		return domain.Person{}, err
	}
	userInfo, err := u.dao.NewUserQuery().GetUserByID(userID, db)
	if err != nil {
		log.Fatal(err)
		return domain.Person{}, err
	}

	return userInfo, nil
}
