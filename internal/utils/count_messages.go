package utils

import (
	"fmt"
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
)

func MidnightRange(loc *time.Location, t time.Time) (time.Time, time.Time) {
	year, month, day := t.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, loc)
	end := start.Add(24 * time.Hour)
	return start, end
}

func CountMessagebyID(userID uint) (int, error) {
	db := DatabaseInstance{}.Instance()
	start, end := MidnightRange(time.Local, time.Now())
	var used int64
	if err := db.Model(&types.Message{}).
		Where("user_id = ? AND created_at >= ? AND created_at < ?", userID, start, end).
		Count(&used).Error; err != nil {
		return 400, fmt.Errorf("Error Get Cant Messages")
	}
	return int(used), nil
}
