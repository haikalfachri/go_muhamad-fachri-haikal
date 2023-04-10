package controllers

import (
	"restful_api_testing/middlewares"
	"restful_api_testing/models"
	"restful_api_testing/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserContoller(jwtAuth *middlewares.JWTConfig) UserController {
	return UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (uc *UserController) GetAll(c echo.Context) error {
	users, err := uc.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch all user",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch all user",
		"user":    users,
	})
}

func (uc *UserController) GetById(c echo.Context) error {
	id := c.Param("id")

	user, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch a user",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch a user",
		"user":    user,
	})
}

func (uc *UserController) Create(c echo.Context) error {
	var userInput models.UserInput
	c.Bind(&userInput)

	user, err := uc.service.Create(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to create a user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to create a user",
		"user":    user,
	})
}

func (uc *UserController) Update(c echo.Context) error {
	var userInput models.UserInput
	c.Bind(&userInput)

	id := c.Param("id")
	user, err := uc.service.Update(userInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update user",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update a user",
		"user":    user,
	})
}

func (uc *UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete a user",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete a user",
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput models.UserInput
	c.Bind(&userInput)

	token, err := uc.service.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to login",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"token": token,
	})
}