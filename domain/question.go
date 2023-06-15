package domain

import "time"

type Question struct {
	ID 			 int
	Title 		 string
	Description  string
	CreatedAt 	 *time.Time
	ModifiedAt 	 *time.Time
	Author       Person
	Category     Category
}
