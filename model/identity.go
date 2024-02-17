package model

import (
	"time"

	"github.com/google/uuid"
)

type Identity struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Role        string    `json:"role"`
	Radius      int       `json:"radius"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	PinCode     string    `gorm:"column:pin_code;type:varchar(255)" json:"pin_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
