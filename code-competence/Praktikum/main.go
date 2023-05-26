package main

import (
  	"code_competence/controllers"
	"code_competence/database"
	"code_competence/middlewares"
	"code_competence/routes"

	echo "github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectDB()

	database.MigrateDB(db)

	configLogger := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	configJWT := middlewares.JWTConfig{
		SecretKey		: "secret_key",
		ExpiresDuration	: 1,
	}

	e := echo.New()

	userCtrl := controllers.InitUserContoller(&configJWT)
	itemCtrl := controllers.InitItemContoller(&configJWT)
	categoryCtrl := controllers.InitCategoryContoller(&configJWT)

	routesInit := routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware	:   configJWT.Init(),
		UserController	: 	*userCtrl,
		ItemController:  *itemCtrl,
		CategoryController: *categoryCtrl,
	}

	routesInit.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
	
}