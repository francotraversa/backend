package services

import (
	"testing"
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllMessageUseCaseOk(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	db.AutoMigrate(&types.Message{}, &types.MessageDestination{})

	utils.OverrideDatabaseInstance(db)

	m1 := types.Message{UserID: 1, Content: "msg A", CreatedAt: time.Unix(1, 0)}
	m2 := types.Message{UserID: 2, Content: "msg B", CreatedAt: time.Unix(2, 0)}
	if err := db.Create(&m1).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&m2).Error; err != nil {
		t.Fatal(err)
	}

	if err := db.Create(&types.MessageDestination{MessageID: m1.ID, Service: "slack", Status: "success"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&types.MessageDestination{MessageID: m2.ID, Service: "telegram", Status: "failed"}).Error; err != nil {
		t.Fatal(err)
	}

	msgs, total, err := GetAllMessageGetUse()
	if err != nil {
		t.Fatalf("ListMessagesAdmin: %v", err)
	}

	if total != 2 {
		t.Fatalf("Debe devolver 2, Obtuve %d", total)
	}
	if len(*msgs) != 2 {
		t.Fatalf("Debe devolver 2, Obtuve %d", len(*msgs))
	}
	if (*msgs)[0].ID != m2.ID {
		t.Fatalf("La primera debe ser la mas nueva id=%d, obtuve id=%d", m2.ID, (*msgs)[0].ID)
	}
}
