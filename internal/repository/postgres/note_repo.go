package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"Personal-Notes/internal/entity"
	"Personal-Notes/internal/logging"
	"Personal-Notes/internal/repository"
)

const (
	sqlCreateNote = `
		INSERT INTO notes (owner_id, title, body, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, owner_id, title, body, created_at, updated_at
	`
	sqlGetByIDNote = `
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
	start := time.Now()

	note.CreatedAt = start

	r.logger.Debug("monitor[note]: starting note db insertion",
		logging.NewField("owner_id", note.OwnerID),
		logging.NewField("title", note.Title),
		logging.NewField("body", note.Body),
		logging.NewField("created_at", note.CreatedAt),
	)

	var resp entity.Note

	err := r.db.QueryRow(ctx, sqlCreateNote,
		note.OwnerID, note.Title, note.Body, note.CreatedAt).
		Scan(&resp.ID, &resp.OwnerID, &resp.Title, &resp.Body, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[note]: %e", repository.ErrTimeout),
				logging.NewField("owner_id", note.OwnerID),
				logging.NewField("title", note.Title),
				logging.NewField("operation", "insert"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.Note{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[note]: %e", repository.ErrDB),
			logging.NewField("owner_id", note.OwnerID),
			logging.NewField("title", note.Title),
			logging.NewField("operation", "insert"),
			logging.NewField("duration", time.Since(start)),
		)
		return entity.Note{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[note]: inserted successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("owner_id", resp.OwnerID),
		logging.NewField("title", resp.Title),
	)
	return resp, nil
}

func (r *NoteRepository) GetByID(ctx context.Context, id int, ownerID int) (entity.Note, error) {
	start := time.Now()

	r.logger.Debug("monitor[note]: starting note db get by id",
		logging.NewField("id", id),
		logging.NewField("owner_id", ownerID),
	)

	var resp entity.Note

	err := r.db.QueryRow(ctx, sqlGetByIDNote, id, ownerID).
		Scan(&resp.ID, &resp.OwnerID, &resp.Title, &resp.Body, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[note]: %e", repository.ErrTimeout),
				logging.NewField("id", id),
				logging.NewField("owner_id", ownerID),
				logging.NewField("operation", "get_by_id"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.Note{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[note]: %e", repository.ErrDB),
			logging.NewField("id", id),
			logging.NewField("owner_id", ownerID),
			logging.NewField("operation", "get_by_id"),
			logging.NewField("duration", time.Since(start)),
		)
		return entity.Note{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[note]: got by id successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("owner_id", resp.OwnerID),
		logging.NewField("title", resp.Title),
	)
	return resp, nil
}

func (r *NoteRepository) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	return entity.Note{}, nil
}

func (r *NoteRepository) Delete(ctx context.Context, id int, ownerID int) error {
	return nil
}
