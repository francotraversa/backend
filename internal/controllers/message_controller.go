package controllers

import (
	"net/http"
	"os"

	"github.com/francotraversa/siriusbackend/internal/auth"
	services "github.com/francotraversa/siriusbackend/internal/services/messages"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MessageController(route *echo.Echo) {
	m := route.Group("/message")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	m.Use(echojwt.WithConfig(config))

	m.POST("/send", sendMessageHandle)
}

func sendMessageHandle(c echo.Context) error {
	uid, err := auth.IdFromContext(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
	}
	var message types.MessageRequest
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, "JSON inválido")
	}
	if message.Content == "" || len(message.Services) == 0 {
		return c.JSON(http.StatusBadRequest, "Contenido y Destinos son obligatorios")
	}
	err2 := services.PostMessageUseCase(uid, message)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err2.Error()})
	}
	return c.JSON(http.StatusOK, "Mensaje cargado")
}
