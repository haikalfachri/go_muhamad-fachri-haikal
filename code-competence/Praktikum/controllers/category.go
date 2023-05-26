package controllers

import (
	"code_competence/middlewares"
	"code_competence/models/response"
	"code_competence/models/input"
	"code_competence/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	service services.CategoryService
}

func InitCategoryContoller(jwtAuth *middlewares.JWTConfig) *CategoryController {
	return &CategoryController{
		service: services.InitCategoryService(jwtAuth),
	}
}

func (uc *CategoryController) Create(c echo.Context) error {
	name := c.FormValue("name")

	var categoryInput input.CategoryInput = input.CategoryInput{
		Name: name,
	}

	err := categoryInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	category, err := uc.service.Create(categoryInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "create success",
		Data	:  category,
	})
}

func (uc *CategoryController) GetAll(c echo.Context) error {
	categorys, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all category",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all category",
		Data	:  categorys,
	})
}

func (uc *CategoryController) GetById(c echo.Context) error {
	id := c.Param("id")

	category, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a category by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a category by id",
		Data	:  category,
	})
}

func (uc *CategoryController) Update(c echo.Context) error {
	name := c.FormValue("name")

	var categoryInput input.CategoryInput = input.CategoryInput{
		Name: name,
	}

	err := categoryInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	category, err := uc.service.Update(categoryInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a category",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update a category",
		Data	:  category,
	})
}

func (uc *CategoryController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a category",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to delete a category",
	})
}

