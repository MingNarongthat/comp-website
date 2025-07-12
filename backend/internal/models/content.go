package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Content struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Key         string    `json:"key" gorm:"uniqueIndex;not null"`
	Title       string    `json:"title"`
	Content     string    `json:"content" gorm:"type:text"`
	Type        string    `json:"type" gorm:"default:'text'"` // text, html, markdown, json
	Section     string    `json:"section"`                    // home, about, contact, etc.
	Order       int       `json:"order" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (content *Content) BeforeCreate(tx *gorm.DB) (err error) {
	if content.ID == uuid.Nil {
		content.ID = uuid.New()
	}
	return
}