package types

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;size:120;not null"`
	Password  string `gorm:"size:255;not null"` // hash
	Role      string `gorm:"size:16;default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}
