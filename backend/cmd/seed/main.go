package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"personal-web/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost port=5432 user=postgres password=postgres dbname=personalweb sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Seed content
	seedContent(db)
	log.Println("Content seeding completed!")
}

func seedContent(db *gorm.DB) {
	contents := []models.Content{
		// Home Hero Section
		{
			Key:      "home.hero.title",
			Title:    "Hero Title",
			Content:  "Welcome to My Personal Web",
			Type:     "text",
			Section:  "home",
			Order:    1,
			IsActive: true,
		},
		{
			Key:      "home.hero.description",
			Title:    "Hero Description",
			Content:  "A passionate full-stack developer sharing insights, experiences, and knowledge through modern web technologies and innovative solutions.",
			Type:     "text",
			Section:  "home",
			Order:    2,
			IsActive: true,
		},
		{
			Key:      "home.hero.about_button",
			Title:    "About Button Text",
			Content:  "Learn More About Me",
			Type:     "text",
			Section:  "home",
			Order:    3,
			IsActive: true,
		},
		{
			Key:      "home.hero.articles_button",
			Title:    "Articles Button Text",
			Content:  "Read My Articles",
			Type:     "text",
			Section:  "home",
			Order:    4,
			IsActive: true,
		},

		// Home Intro Section
		{
			Key:      "home.intro.title",
			Title:    "Introduction Title",
			Content:  "Hello, I'm John Doe",
			Type:     "text",
			Section:  "home",
			Order:    5,
			IsActive: true,
		},
		{
			Key:      "home.intro.subtitle",
			Title:    "Introduction Subtitle",
			Content:  "Full-Stack Developer & Tech Enthusiast",
			Type:     "text",
			Section:  "home",
			Order:    6,
			IsActive: true,
		},
		{
			Key:      "home.intro.paragraph1",
			Title:    "Introduction Paragraph 1",
			Content:  "With over 5 years of experience in modern web development, I specialize in creating scalable applications using cutting-edge technologies like Go, Vue.js, and cloud platforms.",
			Type:     "text",
			Section:  "home",
			Order:    7,
			IsActive: true,
		},
		{
			Key:      "home.intro.paragraph2",
			Title:    "Introduction Paragraph 2",
			Content:  "I'm passionate about sharing knowledge, building innovative solutions, and helping fellow developers grow in their careers through practical insights and real-world examples.",
			Type:     "text",
			Section:  "home",
			Order:    8,
			IsActive: true,
		},
		{
			Key:      "home.intro.skills",
			Title:    "Skills List",
			Content:  `["Go", "Vue.js", "PostgreSQL", "Docker", "Kubernetes", "React", "Node.js"]`,
			Type:     "json",
			Section:  "home",
			Order:    9,
			IsActive: true,
		},
		{
			Key:      "home.intro.photo_placeholder",
			Title:    "Photo Placeholder Text",
			Content:  "Developer Photo",
			Type:     "text",
			Section:  "home",
			Order:    10,
			IsActive: true,
		},

		// Home Articles Section
		{
			Key:      "home.articles.title",
			Title:    "Latest Articles Title",
			Content:  "Latest Articles",
			Type:     "text",
			Section:  "home",
			Order:    11,
			IsActive: true,
		},
		{
			Key:      "home.articles.description",
			Title:    "Latest Articles Description",
			Content:  "Insights, tutorials, and thoughts on web development",
			Type:     "text",
			Section:  "home",
			Order:    12,
			IsActive: true,
		},
		{
			Key:      "home.articles.view_all_button",
			Title:    "View All Articles Button",
			Content:  "View All Articles",
			Type:     "text",
			Section:  "home",
			Order:    13,
			IsActive: true,
		},

		// Home Contact Section
		{
			Key:      "home.contact.title",
			Title:    "Contact Section Title",
			Content:  "Let's Work Together",
			Type:     "text",
			Section:  "home",
			Order:    14,
			IsActive: true,
		},
		{
			Key:      "home.contact.description",
			Title:    "Contact Section Description",
			Content:  "Have a project in mind or want to collaborate? I'd love to hear from you and discuss how we can bring your ideas to life.",
			Type:     "text",
			Section:  "home",
			Order:    15,
			IsActive: true,
		},
		{
			Key:      "home.contact.button_contact",
			Title:    "Contact Button Text",
			Content:  "Get In Touch",
			Type:     "text",
			Section:  "home",
			Order:    16,
			IsActive: true,
		},
		{
			Key:      "home.contact.button_email",
			Title:    "Email Button Text",
			Content:  "Send Email",
			Type:     "text",
			Section:  "home",
			Order:    17,
			IsActive: true,
		},

		// About Page
		{
			Key:      "about.title",
			Title:    "About Page Title",
			Content:  "About Me",
			Type:     "text",
			Section:  "about",
			Order:    1,
			IsActive: true,
		},
		{
			Key:      "about.description",
			Title:    "About Page Description",
			Content:  "Learn more about my background, experience, and passion for web development.",
			Type:     "text",
			Section:  "about",
			Order:    2,
			IsActive: true,
		},

		// Contact Page
		{
			Key:      "contact.title",
			Title:    "Contact Page Title",
			Content:  "Get In Touch",
			Type:     "text",
			Section:  "contact",
			Order:    1,
			IsActive: true,
		},
		{
			Key:      "contact.description",
			Title:    "Contact Page Description",
			Content:  "I'd love to hear from you. Send me a message and I'll get back to you as soon as possible.",
			Type:     "text",
			Section:  "contact",
			Order:    2,
			IsActive: true,
		},

		// General Site Content
		{
			Key:      "site.title",
			Title:    "Site Title",
			Content:  "Personal Web",
			Type:     "text",
			Section:  "general",
			Order:    1,
			IsActive: true,
		},
		{
			Key:      "site.footer.copyright",
			Title:    "Footer Copyright",
			Content:  "© 2025 Personal Web. All rights reserved.",
			Type:     "text",
			Section:  "general",
			Order:    2,
			IsActive: true,
		},
		{
			Key:      "site.email",
			Title:    "Contact Email",
			Content:  "john@example.com",
			Type:     "text",
			Section:  "general",
			Order:    3,
			IsActive: true,
		},
		{
			Key:      "site.social.github",
			Title:    "GitHub URL",
			Content:  "https://github.com/johndoe",
			Type:     "text",
			Section:  "general",
			Order:    4,
			IsActive: true,
		},
		{
			Key:      "site.social.linkedin",
			Title:    "LinkedIn URL",
			Content:  "https://linkedin.com/in/johndoe",
			Type:     "text",
			Section:  "general",
			Order:    5,
			IsActive: true,
		},
		{
			Key:      "site.social.twitter",
			Title:    "Twitter URL",
			Content:  "https://twitter.com/johndoe",
			Type:     "text",
			Section:  "general",
			Order:    6,
			IsActive: true,
		},
	}

	for _, content := range contents {
		// Check if content already exists
		var existingContent models.Content
		result := db.Where("key = ?", content.Key).First(&existingContent)
		
		if result.Error == gorm.ErrRecordNotFound {
			// Content doesn't exist, create it
			if err := db.Create(&content).Error; err != nil {
				log.Printf("Error creating content %s: %v", content.Key, err)
			} else {
				log.Printf("Created content: %s", content.Key)
			}
		} else {
			log.Printf("Content already exists: %s", content.Key)
		}
	}
}