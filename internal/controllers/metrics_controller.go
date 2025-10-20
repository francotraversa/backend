package controllers

import (
	"net/http"
	"os"

	"github.com/francotraversa/siriusbackend/internal/auth"
	services "github.com/francotraversa/siriusbackend/internal/services/admin"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MetricsController(route *echo.Echo) {
	a := route.Group("/admin")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	a.Use(echojwt.WithConfig(config))

	a.GET("/metrics", getcanthandler)
	a.GET("/list", getAllMessage)
}

func getAllMessage(c echo.Context) error {
	role, err := auth.RoleFromContext(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
	}
	if role == "admin" {
		messages, err := services.GetAllMessageGetUse()
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Error make metrics")
		}
		return c.JSON(http.StatusOK, &messages)
	} else {
		return c.JSON(http.StatusUnauthorized, "Sin credenciales Necesarias")
	}
}

func getcanthandler(c echo.Context) error {
	role, err := auth.RoleFromContext(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
	}
	if role == "admin" {
		metrics, err := services.GetCantUseCase()
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Error make metrics")
		}
		return c.JSON(http.StatusOK, &metrics)
	} else {
		return c.JSON(http.StatusUnauthorized, "Sin credenciales Necesarias")
	}
}
