package types

import "time"

type RateLimit struct {
	UserID uint      `gorm:"primaryKey"`
	Day    time.Time `gorm:"type:date;primaryKey"` // solo fecha
	Used   int       `gorm:"not null;default:0"`
}
