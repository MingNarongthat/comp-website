package repository

import (
	"personal-web/internal/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) GetAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Where("is_active = ?", true).Order(`"order" ASC, created_at DESC`).Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) GetByID(id string) (*models.Company, error) {
	var company models.Company
	err := r.db.First(&company, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return r.db.Save(company).Error
}

func (r *CompanyRepository) Delete(id string) error {
	return r.db.Delete(&models.Company{}, "id = ?", id).Error
}