package services

import (
	"personal-web/internal/models"
	"personal-web/internal/repository"

	"github.com/google/uuid"
)

type ContentService interface {
	CreateContent(content *models.Content) error
	GetContentByID(id uuid.UUID) (*models.Content, error)
	GetContentByKey(key string) (*models.Content, error)
	GetContentBySection(section string) ([]models.Content, error)
	GetAllContent() ([]models.Content, error)
	UpdateContent(content *models.Content) error
	DeleteContent(id uuid.UUID) error
}

type contentService struct {
	contentRepo repository.ContentRepository
}

func NewContentService(contentRepo repository.ContentRepository) ContentService {
	return &contentService{
		contentRepo: contentRepo,
	}
}

func (s *contentService) CreateContent(content *models.Content) error {
	return s.contentRepo.Create(content)
}

func (s *contentService) GetContentByID(id uuid.UUID) (*models.Content, error) {
	return s.contentRepo.GetByID(id)
}

func (s *contentService) GetContentByKey(key string) (*models.Content, error) {
	return s.contentRepo.GetByKey(key)
}

func (s *contentService) GetContentBySection(section string) ([]models.Content, error) {
	return s.contentRepo.GetBySection(section)
}

func (s *contentService) GetAllContent() ([]models.Content, error) {
	return s.contentRepo.GetAll()
}

func (s *contentService) UpdateContent(content *models.Content) error {
	return s.contentRepo.Update(content)
}

func (s *contentService) DeleteContent(id uuid.UUID) error {
	return s.contentRepo.Delete(id)
}