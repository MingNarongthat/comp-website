package models

import (
	"time"

	"github.com/google/uuid"
)

// Contact represents a contact form submission
type Contact struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null"`
	Subject     string    `json:"subject" gorm:"not null"`
	Message     string    `json:"message" gorm:"type:text;not null"`
	Phone       string    `json:"phone"`
	Company     string    `json:"company"`
	IsProcessed bool      `json:"is_processed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}