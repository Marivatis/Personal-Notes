package postgres

import (
	"Personal-Notes/internal/logging"
	"Personal-Notes/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepository(db *pgxpool.Pool, logger *logging.Logger) *repository.Repository {
	return &repository.Repository{
		Note: NewNoteRepository(db, logger),
	}
}
