package services

import (
	"testing"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegisterUserUseCaseShouldCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "PasswordTest",
		Role:     nil,
	}

	err = RegisterUserUseCase(newUser)
	if err != nil {
		t.Fatalf("Fallo al registrar el Usuario: %s", err)
	}
}
func TestRegisterUserUseCaseShouldValidateEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName",
		Password: "PasswordTest",
		Role:     nil,
	}

	err = RegisterUserUseCase(newUser)
	if err == nil {
		t.Fatalf("No valido el formato del Email")
	}
}
func TestRegisterUserUseCaseShouldValidatePassword(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "12345",
		Role:     nil,
	}

	err = RegisterUserUseCase(newUser)
	if err == nil {
		t.Fatalf("No valido el formato de la Password")
	}
}
