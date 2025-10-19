package utils

import (
	"log"
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var DBInstance = DatabaseInstance{}

func (database DatabaseInstance) NewDataBase() {
	var db *gorm.DB
	dsn := "host=localhost user=postgres password=postgres dbname=appdb_prod port=5433 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"

	for {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Waiting for DB...", err)
		time.Sleep(2 * time.Second)
	}
	log.Println("DB CONNECTED")
	if err := EnsureHardcodedUser(db); err != nil {
		log.Fatalf("seed harcoded user error %s", err)
	}

	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.Message{})
	db.AutoMigrate(&types.MessageDestination{})
	DBInstance.DB = db
}

func (database DatabaseInstance) Instance() *gorm.DB {
	return DBInstance.DB
}
