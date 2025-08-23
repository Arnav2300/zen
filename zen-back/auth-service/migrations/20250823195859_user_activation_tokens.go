package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserActivationTokens, downUserActivationTokens)
}

func upUserActivationTokens(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE IF NOT EXISTS user_activation_tokens (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token_hash VARCHAR(512) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP NOT NULL,
			used_at TIMESTAMP DEFAULT NULL,
			ip_address VARCHAR(45),
			user_agent VARCHAR(1024)
		);
	`)
	return err
}

func downUserActivationTokens(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS user_activation_tokens;
	`)
	return err
}
