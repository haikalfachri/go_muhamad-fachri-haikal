package controllers

import (
	"code_competence/middlewares"
	"code_competence/models/response"
	"code_competence/models/input"
	"code_competence/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserContoller(jwtAuth *middlewares.JWTConfig) *UserController {
	return &UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (uc *UserController) Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var userInput input.UserInput = input.UserInput{
		Name: name,
		Email: email,
		Password: password,
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	user, err := uc.service.Register(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "register success",
		Data	:  user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var userInput input.UserInput = input.UserInput{
		Email: email,
		Password: password,
	}
	
	token, err := uc.service.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "login failed",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "login success",
		Data	:  token,
	})
}
