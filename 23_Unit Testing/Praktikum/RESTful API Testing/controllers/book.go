package controllers

import (
	"restful_api_testing/models"
	"restful_api_testing/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	service services.BookService
}

func InitBookContoller() BookController {
	return BookController{
		service: services.InitBookService(),
	}
}

func (bc *BookController) GetAll(c echo.Context) error {
	books, err := bc.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch all books",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch all books",
		"book":    books,
	})
}

func (bc *BookController) GetById(c echo.Context) error {
	id := c.Param("id")

	book, err := bc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch a book",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch a book",
		"book":    book,
	})
}

func (bc *BookController) Create(c echo.Context) error {
	var bookInput models.BookInput
	c.Bind(&bookInput)

	book, err := bc.service.Create(bookInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to create a book",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to create a book",
		"book":    book,
	})
}

func (bc *BookController) Update(c echo.Context) error {
	var bookInput models.BookInput
	c.Bind(&bookInput)

	id := c.Param("id")
	book, err := bc.service.Update(bookInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update book",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update a book",
		"book":    book,
	})
}

func (bc *BookController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := bc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete a book",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete a book",
	})
}
