package services

import (
	"personal-web/internal/models"
	"personal-web/internal/repository"
	"github.com/google/uuid"
)

type ArticleService struct {
	articleRepo *repository.ArticleRepository
	userRepo    *repository.UserRepository
}

func NewArticleService(articleRepo *repository.ArticleRepository, userRepo *repository.UserRepository) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	return s.articleRepo.Create(article)
}

func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
	return s.articleRepo.FindAll()
}

func (s *ArticleService) GetArticleByID(id string) (*models.Article, error) {
	return s.articleRepo.FindByID(id)
}

func (s *ArticleService) UpdateArticle(article *models.Article) error {
	return s.articleRepo.Update(article)
}

func (s *ArticleService) DeleteArticle(id string) error {
	return s.articleRepo.Delete(id)
}

func (s *ArticleService) GetUserIDByEmail(email string) (uuid.UUID, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.ID, nil
}
