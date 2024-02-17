package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Chatroom struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;"`
	ChatroomName string         `json:"chatroom_name"`
	Identities   pq.StringArray `gorm:"type:text[]" json:"identities"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Chatrooms []Chatroom
