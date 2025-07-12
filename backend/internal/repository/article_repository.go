package repository

import (
	"personal-web/internal/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticleRepository) FindAll() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *ArticleRepository) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := r.db.First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) Update(article *models.Article) error {
	return r.db.Save(article).Error
}

func (r *ArticleRepository) Delete(id string) error {
	return r.db.Delete(&models.Article{}, "id = ?", id).Error
}
