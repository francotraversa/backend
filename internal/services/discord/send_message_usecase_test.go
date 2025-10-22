package services_discord

import "testing"

func TestSendMessageUseCaseOk(t *testing.T) {
	ok := SendMessagesUseCase("Verificar Discord")
	if ok != "failed" {
		t.Fatalf("Error enviando mensaje")
	}
}
