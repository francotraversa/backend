package main

import (
	"github.com/francotraversa/siriusbackend/internal/controllers"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.DatabaseInstance{}.NewDataBase()
	e := echo.New()
	controllers.RegisterHealth(e)
	controllers.AuthController(e)

	e.Logger.Fatal(e.Start(":" + "8181"))

}
