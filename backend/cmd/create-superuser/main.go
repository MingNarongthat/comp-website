package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"personal-web/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Get command line arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run cmd/create-superuser/main.go <email> <password> <name>")
		fmt.Println("Example: go run cmd/create-superuser/main.go admin@example.com secretpassword \"Super Admin\"")
		os.Exit(1)
	}

	email := os.Args[1]
	userPassword := os.Args[2]
	name := os.Args[3]

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := os.Getenv("DB_HOST")
		if host == "" {
			host = "localhost"
		}
		user := os.Getenv("DB_USER")
		if user == "" {
			user = "postgres"
		}
		dbPassword := os.Getenv("DB_PASSWORD")
		if dbPassword == "" {
			dbPassword = "postgres"
		}
		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			dbname = "personalweb"
		}
		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, dbPassword, dbname)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Check if user already exists
	var existingUser models.User
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		log.Fatal("User with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	// Create superuser
	superuser := models.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
		Role:     "superadmin",
	}

	if err := db.Create(&superuser).Error; err != nil {
		log.Fatal("Failed to create superuser:", err)
	}

	fmt.Printf("✅ Superuser created successfully!\n")
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Role: superadmin\n")
	fmt.Printf("ID: %s\n", superuser.ID)
}