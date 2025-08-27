package domain

import (
	"time"

	"github.com/google/uuid"
)

type ActivationToken struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	TokenHash string    `json:"token_hash"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	UsedAt    time.Time `json:"used_at"`
	IPAdd     string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
}
