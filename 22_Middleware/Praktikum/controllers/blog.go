package controllers

import (
	"mvc_middleware/models"
	"mvc_middleware/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	service services.BlogService
}

func InitBlogContoller() BlogController {
	return BlogController{
		service: services.InitBlogService(),
	}
}

func (bc *BlogController) GetAll(c echo.Context) error {
	blogs, err := bc.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch all blogs",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch all blogs",
		"blog":    blogs,
	})
}

func (bc *BlogController) GetById(c echo.Context) error {
	id := c.Param("id")

	blog, err := bc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to fetch a blog",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to fetch a blog",
		"blog":    blog,
	})
}

func (bc *BlogController) Create(c echo.Context) error {
	var blogInput models.BlogInput
	c.Bind(&blogInput)

	blog, err := bc.service.Create(blogInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to create a blog",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to create a blog",
		"blog":    blog,
	})
}

func (bc *BlogController) Update(c echo.Context) error {
	var blogInput models.BlogInput
	c.Bind(&blogInput)

	id := c.Param("id")
	blog, err := bc.service.Update(blogInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update blog",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update a blog",
		"blog":    blog,
	})
}

func (bc *BlogController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := bc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete a blog",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete a blog",
	})
}
