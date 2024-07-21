package model

import (
	"time"

	"github.com/google/uuid"
)

// Message struct to describe message object.
type ChatMessage struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	ChatroomId   uuid.UUID `json:"chatroom_id"`
	IdentityId   uuid.UUID `json:"identity_id"`
	Message      string    `json:"message"`
	IsRead       bool      `json:"is_read"`
	IdentityRole string    `json:"identity_role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ChatMessages []ChatMessage
