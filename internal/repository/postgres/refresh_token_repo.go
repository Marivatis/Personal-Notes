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
