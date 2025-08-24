package postgres

import (
	"auth-service/internal/storage"
	"auth-service/internal/storage/postgres/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepository(db *pgxpool.Pool) *storage.Repository {
	return &storage.Repository{
		User: user.NewUserRepository(db),
	}
}
