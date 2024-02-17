package model

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Colour             string    `json:"colour"`
	Radius             int       `json:"radius"`
	Address            string    `json:"address"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	RegistrationNumber string    `json:"registration_number"`
	CreatedBy          string    `json:"created_by"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
