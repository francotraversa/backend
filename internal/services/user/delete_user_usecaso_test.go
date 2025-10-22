package services

import (
	"testing"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDeleteUserUseCaseShouldDelete(t *testing.T) {
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
	deleteUser := types.DeleteUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
	}

	RegisterUserUseCase(newUser)

	err = DelteUserUseCase(deleteUser)
	if err != nil {
		t.Fatalf("Usuario no eliminado")
	}
}

func TestDeleteUserUseCaseShouldValidateUser(t *testing.T) {
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
	deleteUser := types.DeleteUser{
		Username: "asdasd",
		Email:    "",
	}

	RegisterUserUseCase(newUser)

	err = DelteUserUseCase(deleteUser)
	if err != nil {
		t.Fatalf("Fallo Porque el Usuario no Existe")
	}
}

func TestDeleteUserUseCaseShouldValidateEmail(t *testing.T) {
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
	deleteUser := types.DeleteUser{
		Username: "",
		Email:    "EmailName@admin.com",
	}

	RegisterUserUseCase(newUser)

	err = DelteUserUseCase(deleteUser)
	if err != nil {
		t.Fatalf("Fallo Porque el Usuario no Existe o Email mal formateado")
	}
}
