package repository

import (
	"context"

	"Personal-Notes/internal/entity"
)

type Note interface {
	Create(ctx context.Context, note entity.Note) (entity.Note, error)
	GetByID(ctx context.Context, id int, ownerID int) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id int, ownerID int) error
}

type Repository struct {
	Note
}
