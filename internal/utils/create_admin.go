package utils

import (
	"log"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const hardcodedUsername = "siriusadmin"
const hardcodedPassword = "admin123"

func EnsureHardcodedUser(db *gorm.DB) error {
	u := strings.ToLower(hardcodedUsername)

	var count int64
	if err := db.Model(&types.User{}).
		Where("LOWER(username) = ?", u).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Printf("[seed] usuario '%s' ya existe, no se crea", hardcodedUsername)
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(hardcodedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := types.User{
		Username: hardcodedUsername,
		Password: string(hash),
		Role:     "admin",
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	log.Printf("[seed] usuario '%s' creado con rol=admin", hardcodedUsername)
	return nil
}
