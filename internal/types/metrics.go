package types

import "time"

type AdminMetricsRow struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	TotalSent      int64  `json:"total_sent"`
	TodaySent      int64  `json:"today_sent"`
	RemainingToday int    `json:"remaining_today"`
}
type MessageAdminResponse struct {
	ID        uint      `json:"MessageID"      `
	UserID    uint      `json:"UserID"  `
	Content   string    `json:"Content" `
	Services  string    `json:"Services"`
	Status    string    `json:"Status"  `
	CreatedAt time.Time `json:"CreatedAt"`
}
