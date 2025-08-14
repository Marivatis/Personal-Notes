package repository

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exists")
	ErrDB           = errors.New("database error")
	ErrTimeout      = errors.New("database query timeout")
)
