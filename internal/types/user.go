package types

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;size:120;not null"`
	Email     string `gorm:"uniqueIndex;size:120;not null"`
	Password  string `gorm:"size:255;not null"` // hash
	Role      string `gorm:"size:16;default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Creds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}
type RegisterUser struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Role     *string `json:"role,omitempty"`
}

type UpdateUser struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Role     *string `json:"role,omitempty"`
}

type DeleteUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
