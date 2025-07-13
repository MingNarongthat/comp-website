package services

import (
	"personal-web/internal/models"
	"personal-web/internal/repository"
)

type CompanyService struct {
	companyRepo *repository.CompanyRepository
}

func NewCompanyService(companyRepo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepo: companyRepo}
}

func (s *CompanyService) CreateCompany(company *models.Company) error {
	return s.companyRepo.Create(company)
}

func (s *CompanyService) GetAllCompanies() ([]models.Company, error) {
	return s.companyRepo.GetAll()
}

func (s *CompanyService) GetCompanyByID(id string) (*models.Company, error) {
	return s.companyRepo.GetByID(id)
}

func (s *CompanyService) UpdateCompany(company *models.Company) error {
	return s.companyRepo.Update(company)
}

func (s *CompanyService) DeleteCompany(id string) error {
	return s.companyRepo.Delete(id)
}