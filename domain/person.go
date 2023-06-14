package domain

import "time"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt *time.Time
}
