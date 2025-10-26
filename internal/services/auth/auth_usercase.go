package authenticator

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/francotraversa/siriusbackend/internal/auth"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
)

func AuthUseCase(userCread types.Creds) (*types.TokenResponse, error) {
	if (userCread.Username == "" && userCread.Email == "") || userCread.Password == "" {
		return nil, fmt.Errorf("Parametros insuficientes")
	}

	var User *types.User

	if strings.TrimSpace(userCread.Email) != "" {
		User = utils.FindUserByEmail(strings.ToLower(strings.TrimSpace(userCread.Email)))
	} else {
		User = utils.FindUserByUsername(strings.ToLower(strings.TrimSpace(userCread.Username)))
	}
	err := utils.CheckPassword(User.Password, userCread.Password)
	if err != nil {
		return nil, fmt.Errorf("Los parametros no es correcto")
	}

	token, err := auth.GenerateToken(User.ID, User.Role, os.Getenv("JWT_SECRET"), 15*time.Minute)
	if err != nil {
		return nil, err
	}
	return token, nil
}
