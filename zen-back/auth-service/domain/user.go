package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID `json:"id"`
	Fname               string    `json:"first_name"`
	Lname               string    `json:"last_name"`
	Email               string    `json:"email"`
	EmailVerified       bool      `json:"email_verified"`
	LastLoginAt         time.Time `json:"last_login_at"`
	LastMagicLinkSentAt time.Time `json:"last_magic_link_sent_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
