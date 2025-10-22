package services

import (
	"testing"
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetCantUseCase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	db.AutoMigrate(&types.User{}, &types.Message{})

	utils.OverrideDatabaseInstance(db)

	u1 := types.User{Username: "alice", Email: "a@a.com"}
	u2 := types.User{Username: "bob", Email: "b@b.com"}
	if err := db.Create(&u1).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&u2).Error; err != nil {
		t.Fatal(err)
	}

	start, end := utils.MidnightRange(time.Local, time.Now())

	msgs := []types.Message{
		{UserID: u1.ID, Content: "hoy-1", CreatedAt: start.Add(2 * time.Hour)},
		{UserID: u1.ID, Content: "hoy-2", CreatedAt: start.Add(5 * time.Hour)},
		{UserID: u1.ID, Content: "ayer", CreatedAt: start.Add(-2 * time.Hour)},
		{UserID: u2.ID, Content: "hoy-bob", CreatedAt: end.Add(-1 * time.Hour)},
	}
	if err := db.Create(&msgs).Error; err != nil {
		t.Fatal(err)
	}

	out, err := GetCantUseCase()
	if err != nil {
		t.Fatalf("GetCantUseCase error: %v", err)
	}

	got := map[uint]types.AdminMetricsRow{}
	for _, r := range *out {
		got[r.UserID] = r
	}

	// u1: total=3, today=2, remaining=98
	if got[u1.ID].TotalSent != 3 || got[u1.ID].TodaySent != 2 || got[u1.ID].RemainingToday != 98 {
		t.Fatalf("u1 want total=3,today=2,rem=98; got total=%d,today=%d,rem=%d",
			got[u1.ID].TotalSent, got[u1.ID].TodaySent, got[u1.ID].RemainingToday)
	}
	// u2: total=1, today=1, remaining=99
	if got[u2.ID].TotalSent != 1 || got[u2.ID].TodaySent != 1 || got[u2.ID].RemainingToday != 99 {
		t.Fatalf("u2 want total=1,today=1,rem=99; got total=%d,today=%d,rem=%d",
			got[u2.ID].TotalSent, got[u2.ID].TodaySent, got[u2.ID].RemainingToday)
	}
}
