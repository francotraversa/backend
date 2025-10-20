package services

import (
	"fmt"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func DelteUserUseCase(user types.DeleteUser) error {
	db := utils.DatabaseInstance{}.Instance()
	if err := db.Where("LOWER(username) = ? OR LOWER(email) = ?", strings.ToLower(user.Username), strings.ToLower(user.Email)).Delete(&types.User{}).Error; err != nil {
		return fmt.Errorf("No se pudo eliminar el usuario")
	}
	return nil
}
