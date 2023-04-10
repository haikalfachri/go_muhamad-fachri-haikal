package controllers

import (
	"restful_api_testing/database"
	
	"github.com/labstack/echo/v4"
)

type testCase struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

func InitEcho() *echo.Echo {
	database.InitDB()
	database.InitialMigration()

	e := echo.New()

	return e
}