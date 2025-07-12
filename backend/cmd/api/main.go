package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=postgres port=5432 user=postgres password=postgres dbname=personalweb sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}, &models.Article{}, &models.Content{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	contentRepo := repository.NewContentRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo, userRepo)
	contentService := services.NewContentService(contentRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	articleHandler := handlers.NewArticleHandler(articleService)
	contentHandler := handlers.NewContentHandler(contentService)

	// API routes
	api := r.Group("/api")
	{
		// User routes
						api.POST("/users/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	api.POST("/create-superuser", userHandler.CreateSuperuser)
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Test route works!"})
		})

		// Public Article routes (no authentication needed for reading)
		api.GET("/articles", articleHandler.GetArticles)
		api.GET("/articles/:id", articleHandler.GetArticle)

		// Public Content routes (no authentication needed for reading)
		api.GET("/content", contentHandler.GetContentBySection)
		api.GET("/content/key/:key", contentHandler.GetContentByKey)
		api.GET("/content/:id", contentHandler.GetContent)

		// Protected Article routes (authentication required for writing)
		protectedArticles := api.Group("/articles")
		protectedArticles.Use(middleware.AuthMiddleware())
		protectedArticles.Use(middleware.SuperadminMiddleware())
		{
			protectedArticles.POST("", articleHandler.CreateArticle)
			protectedArticles.PUT("/:id", articleHandler.UpdateArticle)
			protectedArticles.DELETE("/:id", articleHandler.DeleteArticle)
		}

		// Protected Content routes (authentication required for writing)
		protectedContent := api.Group("/content")
		protectedContent.Use(middleware.AuthMiddleware())
		protectedContent.Use(middleware.SuperadminMiddleware())
		{
			protectedContent.POST("", contentHandler.CreateContent)
			protectedContent.PUT("/:id", contentHandler.UpdateContent)
			protectedContent.DELETE("/:id", contentHandler.DeleteContent)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}