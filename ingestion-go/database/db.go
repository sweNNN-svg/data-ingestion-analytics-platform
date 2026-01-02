package database

import (
	"fmt"
	"log"
	"os"

	"ingestion-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB establishes a connection to PostgreSQL database using GORM
// and automatically migrates the database tables if they don't exist.
func ConnectDB() error {
	// PostgreSQL connection string
	// Format: host=HOST user=USER password=PASSWORD dbname=DBNAME sslmode=SSLMODE
	// Default values match the Python FastAPI service configuration
	host := getEnv("DB_HOST", "db")
	user := getEnv("DB_USER", "user")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "events_db")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	// Open database connection with GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// AutoMigrate automatically creates tables based on the model structs
	// It will create tables if they don't exist, and add missing columns
	// It won't delete unused columns or modify existing columns
	err = DB.AutoMigrate(
		&models.RawEvent{},
		&models.AnalyticsEvent{},
	)

	if err != nil {
		return fmt.Errorf("failed to auto migrate database: %w", err)
	}

	log.Println("Database tables migrated successfully")

	return nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}


