package types

import (
	"time"
)

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageDestination struct {
	ID               uint   `gorm:"primaryKey"`
	MessageID        uint   `gorm:"index"`
	Service          string `gorm:"size:32;index"`
	Status           string `gorm:"size:16;index"` // pending|success|failed
	ProviderResponse string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type MessageRequest struct {
	Content  string    `json:"content"`
	Services []Service `json:"services"`
}

type Service struct {
	App string `json:"app"`
}

type MessageSlack struct {
	Text string `json:"text"`
}
type MessageDiscord struct {
	Content string `json:"content"`
}
