package models

import (
	"time"

	"github.com/google/uuid"
)

// Company represents a trusted company/client
type Company struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	LogoURL     string    `json:"logo_url" gorm:"not null"`
	Website     string    `json:"website"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	Order       int       `json:"order" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}