package main

import (
  "restful_api_testing/database"
  "restful_api_testing/routes"

  "github.com/labstack/echo/v4"
)

func main(){
  database.InitDB()

  database.InitialMigration()

  e := echo.New()
  routes.Setup(e)
  e.Logger.Fatal(e.Start(":8000"))
}
