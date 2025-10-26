package services

import (
	"fmt"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func UpdateUserUseCase(user types.UpdateUser) error {
	db := utils.DatabaseInstance{}.Instance()
	if user.Role == nil || (*user.Role != "admin" && *user.Role != "user") {
		return fmt.Errorf("invalid role")
	}
	res := db.Model(&types.User{}).Where("username = ? OR email = ?", strings.ToLower(user.Username), strings.ToLower(user.Email)).Update("role", *user.Role)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("Usuario no encontrado")
	}
	return nil
}
