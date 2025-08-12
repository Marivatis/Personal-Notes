package entity

import "time"

type User struct {
	id          int
	name        string
	email       string
	password    string
	createdAt   time.Time
	updatedAt   time.Time
	lastLoginAt time.Time
}
