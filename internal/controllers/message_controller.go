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
	m.GET("/today/", getAllMyMessageHandle)
	m.GET("/date/", getMessageByDate)
}

func getMessageByDate(c echo.Context) error {
	message, err := services.GetMessageByDateUseCase(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, message)
}

func getAllMyMessageHandle(c echo.Context) error {
	message, err := services.GetMessageFilterUseCase(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, message)
}

func sendMessageHandle(c echo.Context) error {
	uid, err := auth.IdFromContext(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
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
		return c.JSON(http.StatusTooManyRequests, err2)
	}
	return c.JSON(http.StatusOK, "Mensaje Enviado")
}
