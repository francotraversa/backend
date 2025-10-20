package services

import (
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func GetCantUseCase() (*[]types.AdminMetricsRow, error) {
	db := utils.DatabaseInstance{}.Instance()
	start, end := utils.MidnightRange(time.Local, time.Now())
	totalSub := db.Model(&types.Message{}).
		Select("user_id, COUNT(*) AS total").
		Group("user_id")

	todaySub := db.Model(&types.Message{}).
		Select("user_id, COUNT(*) AS today").
		Where("created_at >= ? AND created_at < ?", start, end).
		Group("user_id")

	var rows []struct {
		UserID    uint
		Username  string
		TotalSent int64
		TodaySent int64
	}
	err := db.Model(&types.User{}).
		Select(`
            users.id AS user_id,
            users.username,
            COALESCE(t_all.total, 0)  AS total_sent,
            COALESCE(t_day.today, 0)  AS today_sent`).
		Joins("LEFT JOIN (?) AS t_all ON t_all.user_id = users.id", totalSub).
		Joins("LEFT JOIN (?) AS t_day ON t_day.user_id = users.id", todaySub).
		Order("users.id").
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}
	out := make([]types.AdminMetricsRow, 0, len(rows))
	for _, r := range rows {
		rem := 100 - int(r.TodaySent)
		if rem < 0 {
			rem = 0
		}
		out = append(out, types.AdminMetricsRow{
			UserID:         r.UserID,
			Username:       r.Username,
			TotalSent:      r.TotalSent,
			TodaySent:      r.TodaySent,
			RemainingToday: rem,
		})
	}
	return &out, nil
}
