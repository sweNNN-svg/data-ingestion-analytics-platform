package main

import (
	"log"
	"os"

	"ingestion-go/database"
	"ingestion-go/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to PostgreSQL database
	// ConnectDB function establishes connection and runs AutoMigrate
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a new Fiber application instance
	// Fiber is a fast HTTP web framework inspired by Express.js
	app := fiber.New(fiber.Config{
		AppName: "Scate Ingestion API",
	})

	// Add CORS middleware
	// CORS (Cross-Origin Resource Sharing) allows web pages to make requests
	// to a different domain than the one serving the web page
	// allow_origins=["*"] means allow requests from any origin
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                    // Her yerden gelen isteğe izin ver
		AllowCredentials: false,                   // Credentials (cookies, auth headers) gönderilmesine izin ver
		AllowMethods:     "GET,POST,PUT,DELETE", // İzin verilen HTTP metodları
		AllowHeaders:     "*",                    // Tüm header'lara izin ver
	}))

	// Define API routes
	// POST /api/events - Event ingestion endpoint
	// Handles incoming event data and saves it to the database
	app.Post("/api/events", handlers.IngestEvent)

	// GET /health - Health check endpoint
	// Simple endpoint to verify that the service is running
	app.Get("/health", handlers.HealthCheck)

	// Get port from environment variable or use default 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the Fiber application server
	// Listen starts the HTTP server on the specified address
	// The server will block until it's stopped (e.g., by Ctrl+C or SIGTERM)
	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// defer anahtar kelimesi hakkında:
	// defer, fonksiyon return olmadan önce çalıştırılacak kodları belirler.
	// Örneğin, bir dosya açtıktan sonra kapatmak için:
	//   file, err := os.Open("file.txt")
	//   defer file.Close()  // Fonksiyon bitmeden önce dosya kapatılacak
	//
	// Burada database.DB için defer kullanmıyoruz çünkü:
	// 1. GORM connection pool'u otomatik yönetir
	// 2. Uygulama çalıştığı sürece bağlantı açık kalmalı
	// 3. Uygulama kapanırken Go runtime otomatik olarak temizlik yapar
	//
	// Ancak eğer manuel olarak bağlantıyı kapatmak isteseydik:
	//   sqlDB, _ := database.DB.DB()
	//   defer sqlDB.Close()
}


