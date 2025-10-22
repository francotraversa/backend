package main

import (
	"github.com/francotraversa/siriusbackend/internal/controllers"
	enviroment "github.com/francotraversa/siriusbackend/internal/enviorement"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	enviroment.LoadEnviroment("dev")
	utils.DatabaseInstance{}.NewDataBase()
	e := echo.New()
	controllers.RegisterHealth(e)
	controllers.AuthController(e)
	controllers.RegisterUserController(e)
	controllers.MetricsController(e)
	controllers.MessageController(e)
	e.Logger.Fatal(e.Start(":" + "8181"))
}
