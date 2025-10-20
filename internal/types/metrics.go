package types

type AdminMetricsRow struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	TotalSent      int64  `json:"total_sent"`
	TodaySent      int64  `json:"today_sent"`
	RemainingToday int    `json:"remaining_today"`
}
