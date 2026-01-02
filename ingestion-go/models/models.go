package models

import (
	"time"

	"gorm.io/datatypes"
)

// RawEvent represents the raw_events table in the database
// JSON tag'leri (json:"...") API'den gelen/giden JSON verilerini Go struct alanlarına map eder.
// Örneğin, JSON'da "eventType" olarak gelen veri, Go'da EventType alanına otomatik olarak atanır.
// Bu sayede API isteklerinde camelCase kullanılabilirken, Go kodunda PascalCase kullanabiliriz.
type RawEvent struct {
	ID            uint           `gorm:"primaryKey;index" json:"id"`
	EventType     string         `gorm:"type:varchar(255);index" json:"eventType"`     // JSON'da "eventType" olarak görünür
	UserID        int            `gorm:"index" json:"userId"`                           // JSON'da "userId" olarak görünür
	Timestamp     time.Time      `gorm:"type:timestamptz" json:"timeStamp"`            // JSON'da "timeStamp" olarak görünür
	MetadataField datatypes.JSON `gorm:"type:jsonb" json:"metaData"`                   // JSON'da "metaData" olarak görünür
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`             // JSON'da "createdAt" olarak görünür
	Source        string         `gorm:"type:varchar(255);index" json:"source"`        // JSON'da "source" olarak görünür
}

// TableName specifies the table name for RawEvent
func (RawEvent) TableName() string {
	return "raw_events"
}

// AnalyticsEvent represents the analytics_events table in the database
// JSON tag'leri, analytics endpoint'inden dönen verilerin JSON formatında nasıl görüneceğini belirler.
type AnalyticsEvent struct {
	ID          uint      `gorm:"primaryKey;index" json:"id"`
	EventType   string    `gorm:"type:varchar(255)" json:"eventType"`   // JSON'da "eventType" olarak görünür
	EventCount  int       `gorm:"type:integer" json:"eventCount"`      // JSON'da "eventCount" olarak görünür
	WindowStart time.Time `gorm:"type:timestamptz" json:"windowStart"` // JSON'da "windowStart" olarak görünür
}

// TableName specifies the table name for AnalyticsEvent
func (AnalyticsEvent) TableName() string {
	return "analytics_events"
}


