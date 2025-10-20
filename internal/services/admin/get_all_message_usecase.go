package services

import (
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func GetAllMessageGetUse() (*[]types.Message, error) {
	db := utils.DatabaseInstance{}.Instance()
	var total int64
	if err := db.Model(&types.Message{}).Count(&total).Error; err != nil {
		return nil, err
	}
	var msgs []types.Message
	if err := db.
		Model(&types.Message{}).
		Select("id, user_id, content, created_at, updated_at").
		Order("messages.created_at DESC").
		Find(&msgs).Error; err != nil {
		return nil, err
	}
	return &msgs, nil

}
