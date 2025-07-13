package services

import (
	"fmt"
	"net/smtp"
	"os"
	"personal-web/internal/models"
	"personal-web/internal/repository"
	"strings"
)

type ContactService struct {
	contactRepo *repository.ContactRepository
}

func NewContactService(contactRepo *repository.ContactRepository) *ContactService {
	return &ContactService{contactRepo: contactRepo}
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	// Save to database
	if err := s.contactRepo.Create(contact); err != nil {
		return err
	}

	// Send email notification
	if err := s.sendEmailNotification(contact); err != nil {
		// Log error but don't fail the contact creation
		fmt.Printf("Failed to send email notification: %v\n", err)
	}

	return nil
}

func (s *ContactService) GetAllContacts() ([]models.Contact, error) {
	return s.contactRepo.GetAll()
}

func (s *ContactService) GetContactByID(id string) (*models.Contact, error) {
	return s.contactRepo.GetByID(id)
}

func (s *ContactService) UpdateContact(contact *models.Contact) error {
	return s.contactRepo.Update(contact)
}

func (s *ContactService) DeleteContact(id string) error {
	return s.contactRepo.Delete(id)
}

func (s *ContactService) sendEmailNotification(contact *models.Contact) error {
	// Get email configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromEmail := os.Getenv("FROM_EMAIL")
	toEmails := os.Getenv("CONTACT_EMAILS") // Comma-separated list

	// Skip email if not configured
	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || fromEmail == "" || toEmails == "" {
		fmt.Println("Email configuration not found, skipping email notification")
		return nil
	}

	// Parse recipient emails
	recipients := strings.Split(toEmails, ",")
	for i, email := range recipients {
		recipients[i] = strings.TrimSpace(email)
	}

	// Create email content
	subject := fmt.Sprintf("New Contact Form Submission: %s", contact.Subject)
	body := fmt.Sprintf(`
New contact form submission received:

Name: %s
Email: %s
Phone: %s
Company: %s
Subject: %s

Message:
%s

Submitted at: %s
`, contact.Name, contact.Email, contact.Phone, contact.Company, contact.Subject, contact.Message, contact.CreatedAt.Format("2006-01-02 15:04:05"))

	// Create email message
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		fromEmail,
		strings.Join(recipients, ","),
		subject,
		body)

	// Send email
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	return smtp.SendMail(addr, auth, fromEmail, recipients, []byte(message))
}