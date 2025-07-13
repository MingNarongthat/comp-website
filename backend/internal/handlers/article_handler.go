package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"personal-web/internal/models"
	"personal-web/internal/services"
)

type ArticleHandler struct {
	articleService *services.ArticleService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var reqBody struct {
		Title   string   `json:"title" binding:"required"`
		Content string   `json:"content" binding:"required"`
		Images  []string `json:"images"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user email from JWT token set by auth middleware
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get user ID from email
	userID, err := h.articleService.GetUserIDByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Convert images slice to JSON
	var imagesJSON []byte
	if len(reqBody.Images) > 0 {
		imagesJSON, _ = json.Marshal(reqBody.Images)
	} else {
		imagesJSON, _ = json.Marshal([]string{})
	}

	article := models.Article{
		Title:   reqBody.Title,
		Content: reqBody.Content,
		Images:  imagesJSON,
		UserID:  userID,
	}

	if err := h.articleService.CreateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	articles, err := h.articleService.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := h.articleService.GetArticleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var reqBody struct {
		Title   string   `json:"title" binding:"required"`
		Content string   `json:"content" binding:"required"`
		Images  []string `json:"images"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleToUpdate, err := h.articleService.GetArticleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Convert images slice to JSON
	var imagesJSON []byte
	if len(reqBody.Images) > 0 {
		imagesJSON, _ = json.Marshal(reqBody.Images)
	} else {
		imagesJSON, _ = json.Marshal([]string{})
	}

	articleToUpdate.Title = reqBody.Title
	articleToUpdate.Content = reqBody.Content
	articleToUpdate.Images = imagesJSON

	if err := h.articleService.UpdateArticle(articleToUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleToUpdate)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := h.articleService.DeleteArticle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ArticleHandler) UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images provided"})
		return
	}

	var imageURLs []string
	
	for _, file := range files {
		// Validate file type
		if !isValidImageType(file.Header.Get("Content-Type")) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPEG, PNG, and GIF are allowed"})
			return
		}

		// Validate file size (max 5MB per image)
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large. Maximum 5MB per image"})
			return
		}

		// Generate unique filename
		filename := generateUniqueFilename(file.Filename)
		
		// Save file
		uploadPath := fmt.Sprintf("./uploads/articles/%s", filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		imageURL := fmt.Sprintf("/uploads/articles/%s", filename)
		imageURLs = append(imageURLs, imageURL)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Images uploaded successfully",
		"images":  imageURLs,
	})
}

func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
	}
	return validTypes[contentType]
}

func generateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	timestamp := time.Now().Unix()
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	return fmt.Sprintf("%d_%x%s", timestamp, randomBytes, ext)
}
