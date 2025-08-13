package postgres

import (
	"Personal-Notes/internal/entity"
	"Personal-Notes/internal/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository struct {
	db     *pgxpool.Pool
	logger *logging.Logger
}

func NewNoteRepository(db *pgxpool.Pool, logger *logging.Logger) *NoteRepository {
	return &NoteRepository{
		db:     db,
		logger: logger,
	}
}

func (r *NoteRepository) Create(ctx context.Context, note entity.Note) (entity.Note, error) {
	return entity.Note{}, nil
}
func (r *NoteRepository) GetById(ctx context.Context, id int) (entity.Note, error) {
	return entity.Note{}, nil
}
func (r *NoteRepository) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	return entity.Note{}, nil
}
func (r *NoteRepository) Delete(ctx context.Context, id int) error {
	return nil
}
