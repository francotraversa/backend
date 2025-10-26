package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(b), err
}
func CheckPassword(hash, plain string) error {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) != nil {
		return fmt.Errorf("Contraseña Incorrecta")
	}
	return nil
}
