package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the model for a user
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Role      string    `json:"role" gorm:"default:'user';not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}