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
	if err := db.Model(&types.User{}).Where("username = ? OR email = ?", strings.ToLower(user.Username), strings.ToLower(user.Email)).Update("role", *user.Role).Error; err != nil {
		return fmt.Errorf("Error al actualizar Role")
	}
	return nil
}
