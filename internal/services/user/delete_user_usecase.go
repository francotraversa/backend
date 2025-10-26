package services

import (
	"fmt"
	"strings"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"gorm.io/gorm"
)

func DelteUserUseCase(user types.DeleteUser) error {
	db := utils.DatabaseInstance{}.Instance()
	var res *gorm.DB
	if user.Email == "" {
		res = db.Where(
			"LOWER(username) = ?",
			strings.ToLower(user.Username),
		).Delete(&types.User{})
	} else if user.Username == "" {
		res = db.Where(
			"LOWER(email) = ?",
			strings.ToLower(user.Email),
		).Delete(&types.User{})
	} else {
		res = db.Where(
			"LOWER(username) = ? AND LOWER(email) = ?",
			strings.ToLower(user.Username),
			strings.ToLower(user.Email),
		).Delete(&types.User{})
	}

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("Usuario no encontrado")
	}
	return nil
}
