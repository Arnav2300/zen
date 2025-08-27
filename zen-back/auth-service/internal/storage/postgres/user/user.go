package user

import (
	"auth-service/internal/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func (r *UserRepository) FindAll(ctx context.Context, page int) ([]*domain.User, int, error) {
	limit := 10
	offset := limit * (page - 1)
	rows, err := r.db.Query(ctx, `
		SELECT count(*) OVER() AS total, id, first_name, last_name, email, email_verified, last_login_at, last_magic_link_sent_at, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*domain.User
	var total int
	for rows.Next() {
		var u domain.User
		var rowTotal int
		err := rows.Scan(
			&rowTotal,
			&u.ID,
			&u.Fname,
			&u.Lname,
			&u.Email,
			&u.EmailVerified,
			&u.LastLoginAt,
			&u.LastMagicLinkSentAt,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			total = rowTotal
		}
		users = append(users, &u)
	}
	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}
	return users, total, nil

}
func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
        SELECT id, first_name, last_name, email, email_verified, last_login_at, last_magic_link_sent_at, created_at, updated_at
        FROM users
        WHERE id = $1
        LIMIT 1
    `, id)

	var u domain.User
	err := row.Scan(
		&u.ID,
		&u.Fname,
		&u.Lname,
		&u.Email,
		&u.EmailVerified,
		&u.LastLoginAt,
		&u.LastMagicLinkSentAt,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
        SELECT id, first_name, last_name, email, email_verified, last_login_at, last_magic_link_sent_at, created_at, updated_at
        FROM users
        WHERE email = $1
        LIMIT 1
    `, email)

	var u domain.User
	err := row.Scan(
		&u.ID,
		&u.Fname,
		&u.Lname,
		&u.Email,
		&u.EmailVerified,
		&u.LastLoginAt,
		&u.LastMagicLinkSentAt,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *UserRepository) SaveUser(ctx context.Context, user domain.User) error {
	_, err := r.db.Exec(ctx, `INSERT INTO users (
            id, first_name, last_name, email, email_verified, last_login_at, last_magic_link_sent_at, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9
        )
    `, user.ID, user.Fname, user.Lname, user.Email, user.EmailVerified, user.LastLoginAt, user.LastMagicLinkSentAt, user.CreatedAt, user.UpdatedAt)
	return err
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
