package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"personal-web/internal/models"
	"personal-web/internal/services"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func NewCompanyHandler(companyService *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var reqBody struct {
		Name        string `json:"name" binding:"required"`
		LogoURL     string `json:"logo_url" binding:"required"`
		Website     string `json:"website"`
		Description string `json:"description"`
		IsActive    *bool  `json:"is_active"`
		Order       *int   `json:"order"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := models.Company{
		Name:        reqBody.Name,
		LogoURL:     reqBody.LogoURL,
		Website:     reqBody.Website,
		Description: reqBody.Description,
		IsActive:    true,
		Order:       0,
	}

	if reqBody.IsActive != nil {
		company.IsActive = *reqBody.IsActive
	}
	if reqBody.Order != nil {
		company.Order = *reqBody.Order
	}

	if err := h.companyService.CreateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
}

func (h *CompanyHandler) GetCompanies(c *gin.Context) {
	companies, err := h.companyService.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	company, err := h.companyService.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	
	company, err := h.companyService.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	var reqBody struct {
		Name        string `json:"name"`
		LogoURL     string `json:"logo_url"`
		Website     string `json:"website"`
		Description string `json:"description"`
		IsActive    *bool  `json:"is_active"`
		Order       *int   `json:"order"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqBody.Name != "" {
		company.Name = reqBody.Name
	}
	if reqBody.LogoURL != "" {
		company.LogoURL = reqBody.LogoURL
	}
	if reqBody.Website != "" {
		company.Website = reqBody.Website
	}
	if reqBody.Description != "" {
		company.Description = reqBody.Description
	}
	if reqBody.IsActive != nil {
		company.IsActive = *reqBody.IsActive
	}
	if reqBody.Order != nil {
		company.Order = *reqBody.Order
	}

	if err := h.companyService.UpdateCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.companyService.DeleteCompany(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

func (h *CompanyHandler) UploadLogo(c *gin.Context) {
	// Get the uploaded file
	file, header, err := c.Request.FormFile("logo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file type
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File must be an image"})
		return
	}

	// Validate file size (max 5MB)
	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size must be less than 5MB"})
		return
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	
	// Create uploads directory if it doesn't exist
	uploadDir := "./uploads/companies"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save file
	filePath := filepath.Join(uploadDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return the file path
	logoURL := fmt.Sprintf("/uploads/companies/%s", filename)
	c.JSON(http.StatusOK, gin.H{
		"logo_url": logoURL,
		"filename": filename,
	})
}