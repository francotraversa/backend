package services

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserUseCase(user types.RegisterUser) error {
	db := utils.DatabaseInstance{}.Instance()
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("Username, Email y Password se encuentran vacias")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return fmt.Errorf("Formato Email invalido")
	}

	if len(user.Password) < 6 {
		return fmt.Errorf("password demasiado corta (min 6)")
	}

	role := "user"

	if user.Role != nil {
		r := strings.ToLower(strings.TrimSpace(*user.Role))
		if r == "user" || r == "admin" {
			role = r
		} else {
			return fmt.Errorf("Rol no permitido")
		}
	}
	var count int64
	if err := db.Model(&types.User{}).Where("LOWER(username) = ? OR LOWER(email) = ?", strings.ToLower(user.Username), user.Email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("El usuario ya existe")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	u := types.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hash),
		Role:     role,
	}
	if err := db.Create(&u).Error; err != nil {
		return fmt.Errorf("Error creando usuario")
	}
	return nil

}
