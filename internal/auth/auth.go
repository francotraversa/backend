package auth

import (
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role, secret string, ttl time.Duration) (*types.TokenResponse, error) {
	exp := time.Now().Add(ttl).Unix()
	claims := &JwtCustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	result := types.TokenResponse{
		Token:   token,
		Expires: exp,
	}
	return &result, nil
}

/*
func ValidateToken(token string, secret string) bool {
	claims := jwt.MapClaims{}
	tok, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodES256 {
			return nil, fmt.Errorf("Algoritmo Invalido")
		}
		return []byte(secret), nil
	})
	if err != nil || !tok.Valid {
		return nil, nil, fmt.Errorf("token inválido: %w", err)
	}
	return tok, claims, nil
}
*/
