package services

import (
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func GetAllMessageGetUse() (*[]types.MessageAdminResponse, int64, error) {
	db := utils.DatabaseInstance{}.Instance()
	var total int64
	if err := db.Model(&types.Message{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var msgs []types.MessageAdminResponse
	err := db.Table("messages").
		Select(`
			messages.id                   AS id,
			messages.user_id              AS user_id,
			messages.content              AS content,
			message_destinations.service  AS services,
			message_destinations.status   AS status,
			messages.created_at           AS created_at`).
		Joins("JOIN message_destinations ON message_destinations.message_id = messages.id").
		Order("messages.created_at DESC, message_destinations.id ASC").
		Scan(&msgs).Error
	if err != nil {
		return nil, 0, err
	}

	return &msgs, total, nil

}
