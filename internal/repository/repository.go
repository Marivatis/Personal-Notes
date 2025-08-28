package repository

import (
	"context"
	"time"

	"Personal-Notes/internal/entity"
)

type Note interface {
	Create(ctx context.Context, note entity.Note) (entity.Note, error)
	GetByID(ctx context.Context, id int, ownerID int) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id int, ownerID int) error
}

type User interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	UpdateLastLoginAt(ctx context.Context, id int, lastLoginAt time.Time) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Note
	User
}
