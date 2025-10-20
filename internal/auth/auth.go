package auth

import (
	"fmt"
	"time"

	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func RoleFromContext(c echo.Context) (string, error) {
	tok, ok := c.Get("user").(*jwt.Token)
	if !ok || tok == nil || !tok.Valid {
		return "", fmt.Errorf("Token Invalido")
	}
	claims, ok := tok.Claims.(*JwtCustomClaims)
	if !ok {
		return "", fmt.Errorf("Error JwtClaims")
	}
	return claims.Role, nil
}

func IdFromContext(c echo.Context) (uint, error) {
	tok, ok := c.Get("user").(*jwt.Token)
	if !ok || tok == nil || !tok.Valid {
		return 0, fmt.Errorf("Token Invalido")
	}
	claims, ok := tok.Claims.(*JwtCustomClaims)
	if !ok {
		return 0, fmt.Errorf("Error JwtClaims")
	}
	return claims.UserId, nil
}
