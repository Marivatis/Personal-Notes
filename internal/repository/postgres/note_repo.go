package postgres

import (
	"Personal-Notes/internal/entity"
	"Personal-Notes/internal/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	sqlCreateNote = `
		INSERT INTO notes (owner_id, title, body, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, owner_id, title, body, created_at, updated_at
	`
	sqlGetByIdNote = `
		SELECT id, owner_id, title, body, created_at, updated_at
		FROM notes
		WHERE id = $1 AND owner_id = $2
	`
	sqlUpdateNote = `
		UPDATE notes
		SET title = $3,
			body = $4,
			updated_at = $5
		WHERE id = $1 AND owner_id = $2 
	`
	sqlDeleteNote = `
		DELETE FROM notes
		WHERE id = $1 AND owner_id = $2
	`
)

type NoteRepository struct {
	db     *pgxpool.Pool
	logger logging.Logger
}

func NewNoteRepository(db *pgxpool.Pool, logger logging.Logger) *NoteRepository {
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
