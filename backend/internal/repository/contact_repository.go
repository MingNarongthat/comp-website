package repository

import (
	"personal-web/internal/models"

	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *ContactRepository) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Order("created_at DESC").Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) GetByID(id string) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.First(&contact, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

func (r *ContactRepository) Delete(id string) error {
	return r.db.Delete(&models.Contact{}, "id = ?", id).Error
}