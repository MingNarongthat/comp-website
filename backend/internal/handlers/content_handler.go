package handlers

import (
	"net/http"
	"personal-web/internal/models"
	"personal-web/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContentHandler struct {
	contentService services.ContentService
}

func NewContentHandler(contentService services.ContentService) *ContentHandler {
	return &ContentHandler{
		contentService: contentService,
	}
}

func (h *ContentHandler) CreateContent(c *gin.Context) {
	var content models.Content
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.contentService.CreateContent(&content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create content"})
		return
	}

	c.JSON(http.StatusCreated, content)
}

func (h *ContentHandler) GetContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	content, err := h.contentService.GetContentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *ContentHandler) GetContentByKey(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content key is required"})
		return
	}

	content, err := h.contentService.GetContentByKey(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *ContentHandler) GetContentBySection(c *gin.Context) {
	section := c.Query("section")
	if section == "" {
		// If no section specified, get all content
		contents, err := h.contentService.GetAllContent()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content"})
			return
		}
		c.JSON(http.StatusOK, contents)
		return
	}

	contents, err := h.contentService.GetContentBySection(section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content"})
		return
	}

	c.JSON(http.StatusOK, contents)
}

func (h *ContentHandler) GetAllContent(c *gin.Context) {
	contents, err := h.contentService.GetAllContent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content"})
		return
	}

	c.JSON(http.StatusOK, contents)
}

func (h *ContentHandler) UpdateContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	var content models.Content
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content.ID = id
	if err := h.contentService.UpdateContent(&content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update content"})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *ContentHandler) DeleteContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	if err := h.contentService.DeleteContent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content deleted successfully"})
}