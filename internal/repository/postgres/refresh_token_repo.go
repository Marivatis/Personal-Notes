package postgres

const (
	sqlCreateRefreshToken = `
		INSERT INTO refresh_tokens (user_id, token, expires_at, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, token, expires_at, created_at, revoked_at, replaced_by_token
	`
	sqlGetByIDRefreshToken = `
		SELECT id, user_id, token, expires_at, created_at, revoked_at, replaced_by_token
		FROM refresh_tokens
		WHERE id = $1 
	`
	sqlGetByHashRefreshToken = `
		SELECT id, user_id, token, expires_at, created_at, revoked_at, replaced_by_token
		FROM refresh_tokens
		WHERE token = $1 
	`
	sqlUpdateRefreshTokenRevokedAt = `
		UPDATE refresh_tokens
		SET revoked_at = $2
		WHERE id = $1
		RETURNING id, revoked_at
	`
	sqlUpdateRefreshTokenReplacedByToken = `
		UPDATE refresh_tokens
		SET replaced_by_token = $2
		WHERE id = $1
		RETURNING id, replaced_by_token
	`
	sqlDeleteRefreshTokenAllExpired = `
		DELETE FROM refresh_tokens
		WHERE expires_at < NOW()
		RETURNING id
	`
	sqlDeleteRefreshTokenAllRevoked = `
		DELETE FROM refresh_tokens
		WHERE revoked_at IS NOT NULL
		RETURNING id
	`
)

type RefreshTokenRepository struct {
	db     *pgxpool.Pool
	logger logging.Logger
}

func NewRefreshTokenRepository(db *pgxpool.Pool, logger logging.Logger) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db:     db,
		logger: logger,
	}
}

func (r *RefreshTokenRepository) Create(
	ctx context.Context,
	refreshToken entity.RefreshToken,
) (entity.RefreshToken, error) {
	start := time.Now()

	refreshToken.CreatedAt = start

	r.logger.Debug("monitor[refresh_token]: starting refresh token db insertion",
		logging.NewField("user_id", refreshToken.UserID),
		logging.NewField("expires_at", refreshToken.ExpiresAt),
		logging.NewField("created_at", refreshToken.CreatedAt),
	)

	var resp entity.RefreshToken

	err := r.db.QueryRow(ctx, sqlCreateRefreshToken,
		refreshToken.UserID, refreshToken.TokenHash, refreshToken.ExpiresAt, refreshToken.CreatedAt).
		Scan(&resp.ID, &resp.UserID, &resp.TokenHash, &resp.ExpiresAt, &resp.CreatedAt, &resp.RevokedAt, &resp.ReplacedByToken)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.logger.Error(fmt.Sprintf("fail[refresh_token]: %v", repository.ErrTimeout),
				logging.NewField("user_id", refreshToken.UserID),
				logging.NewField("operation", "insert"),
				logging.NewField("duration", time.Since(start)),
				logging.NewField("error", err),
			)
			return entity.RefreshToken{}, fmt.Errorf("%w: %w", repository.ErrTimeout, err)
		}
		r.logger.Error(fmt.Sprintf("fail[refresh_token]: %v", repository.ErrDB),
			logging.NewField("user_id", refreshToken.UserID),
			logging.NewField("operation", "insert"),
			logging.NewField("duration", time.Since(start)),
		)
		return entity.RefreshToken{}, fmt.Errorf("%w: %w", repository.ErrDB, err)
	}

	r.logger.Info("done[refresh_token]: inserted successfully",
		logging.NewField("id", resp.ID),
		logging.NewField("user_id", resp.UserID),
	)
	return resp, nil
}

func (r *RefreshTokenRepository) GetByToken(
	ctx context.Context,
	tokenHash string,
) (entity.RefreshToken, error) {
}

func (r *RefreshTokenRepository) RevokeByID(
	ctx context.Context,
	id int,
	revokedAt time.Time,
) error {
}

func (r *RefreshTokenRepository) CleanupExpired(
	ctx context.Context,
) error {
}
