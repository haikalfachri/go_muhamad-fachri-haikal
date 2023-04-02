package routes

import (
	"mvc_middleware/controllers"
	"mvc_middleware/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       "sssstt_rahasia",
		ExpiresDuration: 1,
	}
	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserContoller(&jwtConfig)

	// Endpoint untuk login dan create user tidak memerlukan auth
	skipAuth := e.Group("/skipAuth")
	skipAuth.POST("/users/login", userController.Login)
	skipAuth.POST("/users", userController.Create)

	bookController := controllers.InitBookContoller()

	// Semua endpoint yang memerlukan auth
	useAuth := e.Group("/auth")
	useAuth.Use(echojwt.WithConfig(authMiddlewareConfig))
	useAuth.GET("/users", userController.GetAll)
	useAuth.GET("/users/:id", userController.GetById)
	useAuth.DELETE("/users/:id", userController.Delete)
	useAuth.PUT("/users/:id", userController.Update)
	useAuth.GET("/books", bookController.GetAll)
	useAuth.GET("/books/:id", bookController.GetById)
	useAuth.POST("/books", bookController.Create)
	useAuth.DELETE("/books/:id", bookController.Delete)
	useAuth.PUT("/books/:id", bookController.Update)

	blogController := controllers.InitBlogContoller()
	e.GET("/blogs", blogController.GetAll)
	e.GET("/blogs/:id", blogController.GetById)
	e.POST("/blogs", blogController.Create)
	e.DELETE("/blogs/:id", blogController.Delete)
	e.PUT("/blogs/:id", blogController.Update)
}
