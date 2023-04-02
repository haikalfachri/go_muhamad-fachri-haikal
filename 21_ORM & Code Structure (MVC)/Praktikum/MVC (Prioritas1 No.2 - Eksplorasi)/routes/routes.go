package routes

import (
	"mvc/controllers"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo){
	userController := controllers.InitUserContoller()

	e.GET("/users", userController.GetAll)
	e.GET("/users/:id", userController.GetById)
	e.POST("/users", userController.Create)
	e.DELETE("/users/:id", userController.Delete)
	e.PUT("/users/:id", userController.Update)

	bookController := controllers.InitBookContoller()
	e.GET("/books", bookController.GetAll)
	e.GET("/books/:id", bookController.GetById)
	e.POST("/books", bookController.Create)
	e.DELETE("/books/:id", bookController.Delete)
	e.PUT("/books/:id", bookController.Update)

	blogController := controllers.InitBlogContoller()
	e.GET("/blogs", blogController.GetAll)
	e.GET("/blogs/:id", blogController.GetById)
	e.POST("/blogs", blogController.Create)
	e.DELETE("/blogs/:id", blogController.Delete)
	e.PUT("/blogs/:id", blogController.Update)
}