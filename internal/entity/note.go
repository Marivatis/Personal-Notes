package entity

import "time"

type Note struct {
	ID        int
	OwnerID   int
	Title     string
	Body      *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
