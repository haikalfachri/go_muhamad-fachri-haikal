package routes

import (
	"code_competence/controllers"
	"code_competence/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   		echo.MiddlewareFunc
	JWTMiddleware      		echojwt.Config
	UserController			controllers.UserController
	CategoryController		controllers.CategoryController
	ItemController 			controllers.ItemController
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	noAuth := e.Group("")
	noAuth.POST("/register", cl.UserController.Register)
	noAuth.POST("/login", cl.UserController.Login)

	Auth := e.Group("/auth", echojwt.WithConfig(cl.JWTMiddleware))
	Auth.GET("/items", cl.ItemController.GetAll, middlewares.VerifyToken)
	Auth.GET("/items/:id", cl.ItemController.GetById, middlewares.VerifyToken)
	Auth.POST("/items", cl.ItemController.Create, middlewares.VerifyToken)
	Auth.PUT("/items/:id", cl.ItemController.Update, middlewares.VerifyToken)
	Auth.DELETE("/items/:id", cl.ItemController.Delete, middlewares.VerifyToken)
	Auth.GET("/items/category/:category_id", cl.ItemController.GetItemsByCategoryId, middlewares.VerifyToken)
	Auth.GET("/items", cl.ItemController.GetItemsByName, middlewares.VerifyToken)
	
	Auth.GET("/categories", cl.CategoryController.GetAll, middlewares.VerifyToken)
	Auth.GET("/categories/:id", cl.CategoryController.GetById, middlewares.VerifyToken)
	Auth.POST("/categories", cl.CategoryController.Create, middlewares.VerifyToken)
	Auth.PUT("/categories/:id", cl.CategoryController.Update, middlewares.VerifyToken)
	Auth.DELETE("/categories/:id", cl.CategoryController.Delete, middlewares.VerifyToken)
}
