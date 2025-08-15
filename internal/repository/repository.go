package repository

import (
	"Personal-Notes/internal/entity"
	"context"
)

type Note interface {
	Create(ctx context.Context, note entity.Note) (entity.Note, error)
	GetById(ctx context.Context, id int, ownerId int) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id int, ownerId int) error
}

type Repository struct {
	Note
}
