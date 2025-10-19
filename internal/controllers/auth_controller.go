package controllers

import (
	"encoding/json"
	"io"
	"net/http"
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
	b, errjson := io.ReadAll(c.Request().Body)
	if errjson != nil {
		return c.JSON(http.StatusBadRequest, "Error Read Json Login")
	}

	err := json.Unmarshal(b, &userCread)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error Unmarshal Login")
	}
	if userCread.Username == "" || userCread.Password == "" {
		c.JSON(http.StatusBadRequest, "Todos los campos son requeridos")
	}
	User := utils.FindUserByUsername(userCread.Username)

	token, err := auth.GenerateToken(User.ID, User.Role, "internaltoken", 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Make Token")
	}
	return c.JSON(http.StatusOK, token)
}
