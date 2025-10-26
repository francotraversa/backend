package controllers

import (
	"net/http"

	authenticator "github.com/francotraversa/siriusbackend/internal/services/auth"
	"github.com/francotraversa/siriusbackend/internal/types"
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

	token, err := authenticator.AuthUseCase(userCread)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, token)
}
