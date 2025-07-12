package repository

import (
	"personal-web/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContentRepository interface {
	Create(content *models.Content) error
	GetByID(id uuid.UUID) (*models.Content, error)
	GetByKey(key string) (*models.Content, error)
	GetBySection(section string) ([]models.Content, error)
	GetAll() ([]models.Content, error)
	Update(content *models.Content) error
	Delete(id uuid.UUID) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{db: db}
}

func (r *contentRepository) Create(content *models.Content) error {
	return r.db.Create(content).Error
}

func (r *contentRepository) GetByID(id uuid.UUID) (*models.Content, error) {
	var content models.Content
	err := r.db.Where("id = ?", id).First(&content).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *contentRepository) GetByKey(key string) (*models.Content, error) {
	var content models.Content
	err := r.db.Where("key = ? AND is_active = true", key).First(&content).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *contentRepository) GetBySection(section string) ([]models.Content, error) {
	var contents []models.Content
	err := r.db.Where("section = ? AND is_active = true", section).Order(`"order" ASC, created_at ASC`).Find(&contents).Error
	return contents, err
}

func (r *contentRepository) GetAll() ([]models.Content, error) {
	var contents []models.Content
	err := r.db.Order(`section ASC, "order" ASC, created_at ASC`).Find(&contents).Error
	return contents, err
}

func (r *contentRepository) Update(content *models.Content) error {
	return r.db.Save(content).Error
}

func (r *contentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Content{}, id).Error
}