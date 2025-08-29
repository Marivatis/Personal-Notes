package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"Personal-Notes/internal/entity"
	"Personal-Notes/internal/logging"
	"Personal-Notes/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	sqlCreateUser = `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, password, created_at, updated_at, last_login_at
	`
	sqlGetByIDUser = `
		SELECT id, name, email, password, created_at, updated_at, last_login_at
		FROM users
		WHERE id = $1
	`
	sqlGetByEmailUser = `
		SELECT id, name, email, password, created_at, updated_at, last_login_at
		FROM users
		WHERE email = $1
	`
	sqlUpdateUser = `
		UPDATE users
		SET name = $2,
			 email = $3,
			 password = $4,
			 updated_at = $5
		WHERE id = $1
		RETURNING id, name, email, password, created_at, updated_at, last_login_at
	`
	sqlUpdateUserLastLoginAt = `
		UPDATE users
		SET updated_at = $2,
			 last_login_at = $3
		WHERE id = $1
	`
	sqlDeleteUser = `
		DELETE FROM users
		WHERE id = $1
	`
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger logging.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger logging.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	start := time.Now()

	user.CreatedAt = start

	r.logger.Debug("monitor[user]: starting user db insertion",
		logging.NewField("name", user.Name),
		logging.NewField("email", user.Email),
		logging.NewField("password", user.Password),
		logging.NewField("created_at", user.CreatedAt),
	)

	var resp entity.User

	err := r.db.QueryRow(ctx, sqlCreateUser,
		user.Name, user.Email, user.Password, user.CreatedAt).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.CreatedAt, &resp.UpdatedAt, &resp.LastLoginAt)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("name", user.Name),
				logging.NewField("email", user.Email),
				logging.NewField("operation", "insert"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("name", user.Name),
			logging.NewField("email", user.Email),
			logging.NewField("operation", "insert"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return entity.User{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[user]: inserted successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("name", user.Name),
		logging.NewField("email", user.Email),
	)
	return resp, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (entity.User, error) {
	start := time.Now()

	r.logger.Debug("monitor[user]: starting user db get by id",
		logging.NewField("id", id),
	)

	var resp entity.User

	err := r.db.QueryRow(ctx, sqlGetByIDUser, id).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.CreatedAt, &resp.UpdatedAt, &resp.LastLoginAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrNotFound),
				logging.NewField("id", id),
				logging.NewField("operation", "get_by_id"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrNotFound, err)
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("id", id),
				logging.NewField("operation", "get_by_id"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("id", id),
			logging.NewField("operation", "get_by_id"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return entity.User{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[user]: got by id successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("name", resp.Name),
	)
	return resp, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	start := time.Now()

	r.logger.Debug("monitor[user]: starting user db get by email",
		logging.NewField("email", email),
	)

	var resp entity.User

	err := r.db.QueryRow(ctx, sqlGetByEmailUser, email).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.CreatedAt, &resp.UpdatedAt, &resp.LastLoginAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrNotFound),
				logging.NewField("email", email),
				logging.NewField("operation", "get_by_email"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrNotFound, err)
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("email", email),
				logging.NewField("operation", "get_by_email"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("email", email),
			logging.NewField("operation", "get_by_email"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return entity.User{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[user]: got by email successfully",
		logging.NewField("email", resp.Email),
		logging.NewField("name", resp.Name),
	)
	return resp, nil
}

func (r *UserRepository) Update(ctx context.Context, user entity.User) (entity.User, error) {
	start := time.Now()

	user.UpdatedAt = &start

	r.logger.Debug("monitor[user]: starting user db update",
		logging.NewField("id", user.ID),
		logging.NewField("name", user.Name),
		logging.NewField("email", user.Email),
		logging.NewField("created_at", user.CreatedAt),
		logging.NewField("updated_at", user.UpdatedAt),
		logging.NewField("last_login_at", user.LastLoginAt),
	)

	var resp entity.User

	err := r.db.QueryRow(ctx, sqlUpdateUser, user.ID, user.Name, user.Email, user.Password, user.UpdatedAt).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.CreatedAt, &resp.UpdatedAt, &resp.LastLoginAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrNotFound),
				logging.NewField("id", user.ID),
				logging.NewField("operation", "update"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrNotFound, err)
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("id", user.ID),
				logging.NewField("operation", "update"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.User{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("id", user.ID),
			logging.NewField("operation", "update"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return entity.User{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[user]: updated successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("name", resp.Name),
	)
	return resp, nil
}

func (r *UserRepository) UpdateLastLoginAt(ctx context.Context, id int, lastLoginAt time.Time) error {
	start := time.Now()

	updatedAt := &start

	r.logger.Debug("monitor[user]: starting user lastLoginAt db update",
		logging.NewField("id", id),
		logging.NewField("updated_at", updatedAt),
		logging.NewField("last_login_at", lastLoginAt),
	)

	tag, err := r.db.Exec(ctx, sqlUpdateUserLastLoginAt, id, updatedAt, lastLoginAt)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("id", id),
				logging.NewField("operation", "update_lastLoginAt"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("id", id),
			logging.NewField("operation", "update_lastLoginAt"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return fmt.Errorf("%w: %w", repository.ErrDB, err)
	}
	if tag.RowsAffected() == 0 {
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrNotFound),
			logging.NewField("id", id),
			logging.NewField("operation", "update_lastLoginAt"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return fmt.Errorf("%w: %w", repository.ErrNotFound, err)
	}

	r.logger.Info("done[user]: lastLoginAt updated successfully",
		logging.NewField("id", id),
	)
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	start := time.Now()

	r.logger.Debug("monitor[user]: starting user db delete",
		logging.NewField("id", id),
	)

	tag, err := r.db.Exec(ctx, sqlDeleteUser, id)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrTimeout),
				logging.NewField("id", id),
				logging.NewField("operation", "delete"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrDB),
			logging.NewField("id", id),
			logging.NewField("operation", "delete"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return fmt.Errorf("%w: %w", repository.ErrDB, err)
	}
	if tag.RowsAffected() == 0 {
		r.logger.Error(fmt.Sprintf("fail[user]: %v", repository.ErrNotFound),
			logging.NewField("id", id),
			logging.NewField("operation", "delete"),
			logging.NewField("duration", time.Since(start)),
			logging.NewField("error", err),
		)
		return fmt.Errorf("%w: %w", repository.ErrNotFound, err)
	}

	r.logger.Info("done[user]: deleted successfully",
		logging.NewField("id", id),
	)
	return nil
}
