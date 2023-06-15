package service

import (
	"cserver/domain"
	"cserver/repository"
	"database/sql"
	"log"
)

type CategoryService interface {
	GetCategories(db *sql.DB) ([]domain.Category, error)
}

type categoryService struct {
	dao repository.DAO
}

func NewCategoryQuery(dao repository.DAO) CategoryService {
	return &categoryService{dao: dao}
}

func (c *categoryService) GetCategories(db *sql.DB) ([]domain.Category, error){
	categories, err := c.dao.NewCategoryQuery().GetCategories(db)

	if err != nil {
		log.Fatal(err)
	}

	return categories, nil
}
