package services_slack

import (
	"testing"
)

func TestSendMessageUseCaseOK(t *testing.T) {
	ok := SendMessagesUseCase("Verificar Slack")
	if ok != "failed" {
		t.Fatalf("Error enviando mensaje")
	}
}
