package services_discord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/francotraversa/siriusbackend/internal/types"
)

func SendMessagesUseCase(text string) string {
	body, err := json.Marshal(types.MessageDiscord{Content: text})
	if err != nil {
		return "failed"
	}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("DISCORD_WH"), bytes.NewReader(body))

	if err != nil {
		return "failed"
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "failed"
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "failed"
	}

	return "success"
}
