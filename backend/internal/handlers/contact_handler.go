package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"personal-web/internal/models"
	"personal-web/internal/services"
)

type ContactHandler struct {
	contactService *services.ContactService
}

func NewContactHandler(contactService *services.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var reqBody struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Subject string `json:"subject" binding:"required"`
		Message string `json:"message" binding:"required"`
		Phone   string `json:"phone"`
		Company string `json:"company"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := models.Contact{
		Name:    reqBody.Name,
		Email:   reqBody.Email,
		Subject: reqBody.Subject,
		Message: reqBody.Message,
		Phone:   reqBody.Phone,
		Company: reqBody.Company,
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Thank you for your message! We'll get back to you soon.",
		"id":      contact.ID,
	})
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	contacts, err := h.contactService.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func (h *ContactHandler) GetContact(c *gin.Context) {
	id := c.Param("id")
	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	
	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	var reqBody struct {
		IsProcessed *bool `json:"is_processed"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqBody.IsProcessed != nil {
		contact.IsProcessed = *reqBody.IsProcessed
	}

	if err := h.contactService.UpdateContact(contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.contactService.DeleteContact(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}