package models

import "time"

type User struct {
	ID              string
	FirstName       string
	LastName        string
	Password        string
	PhoneNumber     string
	Email           string
	IsEmailVerified bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
