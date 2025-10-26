package services

import (
	"fmt"
	"strings"

	services_discord "github.com/francotraversa/siriusbackend/internal/services/discord"
	services_slack "github.com/francotraversa/siriusbackend/internal/services/slack"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func PostMessageUseCase(id uint, message types.MessageRequest) error {
	db := utils.DatabaseInstance{}.Instance()
	newmessages := types.Message{
		UserID:  id,
		Content: message.Content,
	}
	cant, err := utils.CountMessagebyID(id)
	if cant <= 3 {
		if err != nil {
			return err
		}

		if err := db.Create(&newmessages).Error; err != nil {
			return fmt.Errorf("Error create row in Messages")
		}

		seen := map[string]struct{}{}
		dests := make([]types.MessageDestination, 0, len(message.Services))

		for _, s := range message.Services {
			app := strings.ToLower(strings.TrimSpace(s.App))
			if app == "" {
				continue
			}
			var statu string
			switch app {
			case "slack":
				statu = services_slack.SendMessagesUseCase(message.Content)
			case "discord":
				statu = services_discord.SendMessagesUseCase(message.Content)
			}
			if _, dup := seen[app]; dup {
				continue
			}
			seen[app] = struct{}{}

			dests = append(dests, types.MessageDestination{
				MessageID: newmessages.ID,
				Service:   app,
				Status:    statu,
			})
		}
		if err := db.Omit("ID").Create(&dests).Error; err != nil {
			return fmt.Errorf("Error create row in MessageDestination")
		}
	} else {
		return fmt.Errorf("Limite de Mensajes alcanzados")
	}
	return nil

}
