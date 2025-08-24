package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"Personal-Notes/internal/logging"
	"Personal-Notes/internal/repository"
)

func NewRepository(db *pgxpool.Pool, logger logging.Logger) *repository.Repository {
	return &repository.Repository{
		Note: NewNoteRepository(db, logger),
	}
}
