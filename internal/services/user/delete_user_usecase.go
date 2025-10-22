package services

import (
	"fmt"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func DelteUserUseCase(user types.DeleteUser) error {
	db := utils.DatabaseInstance{}.Instance()
	res := db.Where(
		"LOWER(username) = ? OR LOWER(email) = ?",
		strings.ToLower(user.Username),
		strings.ToLower(user.Email),
	).Delete(&types.User{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("Usuario no encontrado")
	}
	return nil
}
