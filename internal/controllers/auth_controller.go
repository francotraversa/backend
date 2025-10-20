package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/francotraversa/siriusbackend/internal/auth"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthController(route *echo.Echo) {
	a := route.Group("/auth")
	a.POST("/login", loginHandler)

}

func loginHandler(c echo.Context) error {
	var userCread types.Creds
	if err := c.Bind(&userCread); err != nil {
		return c.JSON(http.StatusBadRequest, "JSON inválido")
	}

	if (userCread.Username == "" && userCread.Email == "") || userCread.Password == "" {
		c.JSON(http.StatusBadRequest, "Todos los campos son requeridos")
	}

	var User *types.User

	if strings.TrimSpace(userCread.Email) != "" {
		User = utils.FindUserByEmail(strings.ToLower(strings.TrimSpace(userCread.Email)))
	} else {
		User = utils.FindUserByUsername(strings.ToLower(strings.TrimSpace(userCread.Username)))
	}
	if User == nil {
		return c.JSON(http.StatusUnauthorized, "Credenciales inválidas")
	}

	token, err := auth.GenerateToken(User.ID, User.Role, os.Getenv("JWT_SECRET"), 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Make Token")
	}
	return c.JSON(http.StatusOK, token)
}
