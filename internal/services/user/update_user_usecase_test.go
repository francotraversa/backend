package services

import (
	"testing"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUpdateUserUseCaseShouldUpdateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	role := "admin"
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "PasswordTest",
		Role:     &role,
	}
	RegisterUserUseCase(newUser)
	role = "user"
	user := types.UpdateUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Role:     &role,
	}
	err = UpdateUserUseCase(user)
	if err == nil {
		t.Fatalf("Usuario no actualizado")
	}
}
func TestUpdateUserUseCaseShouldValidateUsername(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	role := "admin"
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "PasswordTest",
		Role:     &role,
	}
	RegisterUserUseCase(newUser)
	role = "user"
	user := types.UpdateUser{
		Username: "jorge",
		Email:    "",
		Role:     &role,
	}
	err = UpdateUserUseCase(user)
	if err == nil {
		t.Fatalf("Error: No existe el usuario")
	}
}
func TestUpdateUserUseCaseShouldValidateEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	role := "admin"
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "PasswordTest",
		Role:     &role,
	}
	RegisterUserUseCase(newUser)
	role = "user"
	user := types.UpdateUser{
		Username: "",
		Email:    "EmailName@admin.com",
		Role:     &role,
	}
	err = UpdateUserUseCase(user)
	if err == nil {
		t.Fatalf("Error: No existe el usuario con ese Email")
	}
}
func TestUpdateUserUseCaseShouldValidateRole(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&types.User{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	utils.OverrideDatabaseInstance(db)
	role := "admin"
	newUser := types.RegisterUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Password: "PasswordTest",
		Role:     &role,
	}
	RegisterUserUseCase(newUser)
	role = "sinrole"
	user := types.UpdateUser{
		Username: "TestName",
		Email:    "EmailName@admin.com",
		Role:     &role,
	}
	err = UpdateUserUseCase(user)
	if err == nil {
		t.Fatalf("Error: Tiene que ser un rol valido")
	}
}
