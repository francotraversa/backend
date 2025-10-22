package services

import (
	"fmt"
	"time"

	"github.com/francotraversa/siriusbackend/internal/auth"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"github.com/labstack/echo/v4"
)

func GetMessageFilterUseCase(c echo.Context) (*[]types.MessageFilter, error) {
	db := utils.DatabaseInstance{}.Instance()
	uid, err := auth.IdFromContext(c)
	if err != nil {
		return nil, err
	}
	statuses := utils.NormalizarConsulta(c.QueryParam("status"))
	services := utils.NormalizarConsulta(c.QueryParam("service"))
	start, end := utils.MidnightRange(time.Local, time.Now())

	q := db.Model(&types.MessageDestination{}).
		Select(`
			message_destinations.id        AS dest_id,
			message_destinations.service   AS service,
			message_destinations.status    AS status,
			messages.id                    AS message_id,
			messages.content               AS content,
			messages.created_at            AS created_at`).
		Joins("JOIN messages ON messages.id = message_destinations.message_id").
		Where("messages.user_id = ?", uid).
		Where("messages.created_at >= ? AND messages.created_at < ?", start, end)

	if len(statuses) > 0 {
		q = q.Where("message_destinations.status IN ?", statuses)
	}
	if len(services) > 0 {
		q = q.Where("message_destinations.service IN ?", services)
	}

	var mf []types.MessageFilter

	if err := q.Order("messages.created_at DESC, message_destinations.id ASC").
		Limit(100).Offset(0).
		Find(&mf).Error; err != nil {
		return nil, fmt.Errorf("Error data base")
	}
	return &mf, nil

}
