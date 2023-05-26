package controllers

import (
	"code_competence/middlewares"
	"code_competence/models/response"
	"code_competence/models/input"
	"code_competence/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	service services.ItemService
}

func InitItemContoller(jwtAuth *middlewares.JWTConfig) *ItemController {
	return &ItemController{
		service: services.InitItemService(jwtAuth),
	}
}

func (uc *ItemController) Create(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	stock := c.FormValue("stock")
	int_stock, _ := strconv.ParseInt(stock, 10, 32)
	price := c.FormValue("price")
	int_price, _ := strconv.ParseInt(price, 10, 32)
	category_id := c.FormValue("category_id")
	uint_category_id, _ := strconv.ParseUint(category_id, 10, 32)

	var itemInput input.ItemInput = input.ItemInput{
		Name: name,
		Description: description,
		Stock: int(int_stock),
		Price: int(int_price),
		CategoryID: uint(uint_category_id),
	}

	err := itemInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	item, err := uc.service.Create(itemInput)
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
		Data	:  item,
	})
}

func (uc *ItemController) GetAll(c echo.Context) error {
	items, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all item",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all item",
		Data	:  items,
	})
}

func (uc *ItemController) GetById(c echo.Context) error {
	id := c.Param("id")

	item, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch an item by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch an item by id",
		Data	:  item,
	})
}

func (uc *ItemController) GetItemsByCategoryId(c echo.Context) error {
	id := c.Param("category_id")

	items, err := uc.service.GetItemsByCategoryId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all items by category id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a all items by category id",
		Data	:  items,
	})
}

func (uc *ItemController) GetItemsByName(c echo.Context) error {
	name := c.QueryParam("keyword")

	items, err := uc.service.GetItemsByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all items by name",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all items by name",
		Data	:  items,
	})
}

func (uc *ItemController) Update(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	stock := c.FormValue("stock")
	int_stock, _ := strconv.ParseInt(stock, 10, 32)
	price := c.FormValue("price")
	int_price, _ := strconv.ParseInt(price, 10, 32)
	category_id := c.FormValue("category_id")
	uint_category_id, _ := strconv.ParseUint(category_id, 10, 32)

	var itemInput input.ItemInput = input.ItemInput{
		Name: name,
		Description: description,
		Stock: int(int_stock),
		Price: int(int_price),
		CategoryID: uint(uint_category_id),
	}

	err := itemInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	item, err := uc.service.Update(itemInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to update an item",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update an item",
		Data	:  item,
	})
}

func (uc *ItemController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete an item",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to delete an item",
	})
}

