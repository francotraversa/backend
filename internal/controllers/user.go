package controllers

import (
	"net/http"
	"os"

	"github.com/francotraversa/siriusbackend/internal/auth"
	services "github.com/francotraversa/siriusbackend/internal/services/user"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterUserController(route *echo.Echo) {
	u := route.Group("/loged")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	u.Use(echojwt.WithConfig(config))
	u.POST("/reguser", registerhandle)
	u.PATCH("/upduser", upduserhandle)
	u.DELETE("/deluser", deluserhandle)
}

func deluserhandle(c echo.Context) error {
	var delUser types.DeleteUser
	role, errtoken := auth.RoleFromContext(c)
	if errtoken != nil {
		return c.JSON(http.StatusForbidden, errtoken)
	}
	if role == "admin" {
		if err := c.Bind(&delUser); err != nil {
			return c.JSON(http.StatusBadRequest, "JSON inválido")
		}

		err := services.DelteUserUseCase(delUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "Usuario Eliminado Correctamente")
	}

	return c.JSON(http.StatusUnauthorized, "Sin credenciales Necesarias")
}
func registerhandle(c echo.Context) error {
	var newUser types.RegisterUser
	role, errtoken := auth.RoleFromContext(c)
	if errtoken != nil {
		return c.JSON(http.StatusForbidden, errtoken)
	}
	if role == "admin" {
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, "JSON inválido")
		}

		err := services.RegisterUserUseCase(newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "Usuario Creado Correctamente")
	}

	return c.JSON(http.StatusUnauthorized, "Sin credenciales Necesarias")
}

func upduserhandle(c echo.Context) error {
	var updateUser types.UpdateUser
	role, errtoken := auth.RoleFromContext(c)
	if errtoken != nil {
		return c.JSON(http.StatusForbidden, errtoken)
	}
	if role == "admin" {
		if err := c.Bind(&updateUser); err != nil {
			return c.JSON(http.StatusBadRequest, "JSON inválido")
		}
		err := services.UpdateUserUseCase(updateUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "Usuario actualizado correctamente")
	}
	return c.JSON(http.StatusUnauthorized, "Sin credenciales Necesarias")
}
