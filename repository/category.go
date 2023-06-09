package repository

import (
	"cserver/domain"
	"database/sql"
	"log"
)

type CategoryQuery interface{
	GetCategories(db *sql.DB) ([]domain.Category, error)
	GetCategoryByName (name string, db *sql.DB) (domain.Category, error)
}

type categoryQuery struct {}

func (c *categoryQuery) GetCategories(db *sql.DB) ([]domain.Category, error){
	query := `SELECT * FROM "category"`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var categories []domain.Category
	var category domain.Category

	for rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		log.Fatal()
	}

	return categories, nil
}

func (c *categoryQuery) GetCategoryByName (name string, db *sql.DB) (domain.Category, error) {
	query := `SELECT * FROM "category" WHERE "name" = $1`

	var category domain.Category
	err := db.QueryRow(query, name).Scan(
		&category.ID, 
		&category.Name, 
		&category.CreatedAt)
	
	if err != nil {
		log.Fatal(err)
	}

	return category, nil
}
