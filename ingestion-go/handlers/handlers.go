package handlers

import (
	"ingestion-go/database"
	"ingestion-go/models"

	"github.com/gofiber/fiber/v2"
)

// IngestEvent handles POST /events endpoint
// It receives event data in JSON format, validates it, and saves it to the database.
//
// Go'da error handling pattern'i (if err != nil):
// Go, exception-based error handling yerine explicit error return kullanır.
// Her fonksiyon hata durumunda error döndürür ve bu hatalar kontrol edilmelidir.
// Bu yaklaşım:
// 1. Hataların gözden kaçmasını önler (try-catch'te unutulabilir)
// 2. Hangi adımda hata olduğunu açıkça gösterir
// 3. Error handling'i zorunlu kılar (compile-time güvenlik)
// 4. Kod akışını daha okunabilir hale getirir
func IngestEvent(c *fiber.Ctx) error {
	// Create a new RawEvent instance to hold the parsed JSON data
	var event models.RawEvent

	// Parse JSON body into the RawEvent struct
	// BodyParser automatically maps JSON fields to struct fields using json tags
	// Eğer JSON formatı hatalıysa veya gerekli alanlar eksikse hata döner
	if err := c.BodyParser(&event); err != nil {
		// JSON parse hatası durumunda 400 Bad Request döndür
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
			"detail": err.Error(),
		})
	}

	// Save the event to database using GORM
	// Create method inserts a new record into the database
	// Eğer veritabanı bağlantısı yoksa veya constraint hatası varsa hata döner
	if err := database.DB.Create(&event).Error; err != nil {
		// Veritabanı hatası durumunda 500 Internal Server Error döndür
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save event to database",
			"detail": err.Error(),
		})
	}

	// Başarılı kayıt durumunda 201 Created status code ile kaydedilen veriyi döndür
	// GORM, kayıt sonrası event.ID ve event.CreatedAt gibi otomatik alanları doldurur
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event ingested successfully",
		"event":   event,
	})
}

// HealthCheck handles GET /health endpoint
// Simple health check endpoint that returns "OK" status
// Used to verify that the service is running and responsive
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}


