package controllers

import (
	"net/http"
	"os"

	"github.com/francotraversa/siriusbackend/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterHealth(route *echo.Echo) {
	h := route.Group("/health")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	h.Use(echojwt.WithConfig(config))
	h.GET("", getstatushandler)
}

func getstatushandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
