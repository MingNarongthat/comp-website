package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"personal-web/config"
	"personal-web/internal/handlers"
	"personal-web/internal/middleware"
	"personal-web/internal/models"
	"personal-web/internal/repository"
	"personal-web/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}, &models.Article{}, &models.Company{}, &models.Contact{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	companyRepo := repository.NewCompanyRepository(db)
	contactRepo := repository.NewContactRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo, userRepo)
	companyService := services.NewCompanyService(companyRepo)
	contactService := services.NewContactService(contactRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	articleHandler := handlers.NewArticleHandler(articleService)
	companyHandler := handlers.NewCompanyHandler(companyService)
	contactHandler := handlers.NewContactHandler(contactService)

	// Initialize router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Serve static files
	r.Static("/uploads", "./uploads")

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Public routes
		api.POST("/users/register", userHandler.Register)
		api.POST("/users/login", userHandler.Login)
		api.GET("/articles", articleHandler.GetArticles)
		api.GET("/articles/:id", articleHandler.GetArticle)
		api.GET("/companies", companyHandler.GetCompanies)
		api.POST("/contact", contactHandler.CreateContact)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/users/:id", userHandler.GetUser)
			protected.POST("/articles", articleHandler.CreateArticle)
			protected.PUT("/articles/:id", articleHandler.UpdateArticle)
			protected.DELETE("/articles/:id", articleHandler.DeleteArticle)
			protected.POST("/articles/upload-images", articleHandler.UploadImages)
			protected.POST("/companies", companyHandler.CreateCompany)
			protected.GET("/companies/:id", companyHandler.GetCompany)
			protected.PUT("/companies/:id", companyHandler.UpdateCompany)
			protected.DELETE("/companies/:id", companyHandler.DeleteCompany)
			protected.POST("/companies/upload-logo", companyHandler.UploadLogo)
			protected.GET("/contacts", contactHandler.GetContacts)
			protected.GET("/contacts/:id", contactHandler.GetContact)
			protected.PUT("/contacts/:id", contactHandler.UpdateContact)
			protected.DELETE("/contacts/:id", contactHandler.DeleteContact)
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Server failed to start:", err)
	}
} 