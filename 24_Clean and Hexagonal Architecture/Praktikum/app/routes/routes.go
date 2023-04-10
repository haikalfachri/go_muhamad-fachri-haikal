package routes

import (
	"belajar-go-echo/controllers/users"
	"belajar-go-echo/app/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	AuthController     users.AuthController
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	useAuth := e.Group("/auth/users", echojwt.WithConfig(cl.JWTMiddleware))
	noAuth := e.Group("/users")

	useAuth.Use(middlewares.VerifyToken)
	useAuth.GET("", cl.AuthController.GetAllUsers)
	noAuth.POST("/login", cl.AuthController.Login)
	noAuth.POST("", cl.AuthController.CreateUser)
}