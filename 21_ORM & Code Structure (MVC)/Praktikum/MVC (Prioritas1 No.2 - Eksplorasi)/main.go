package main

import (
  "mvc/database"
  "mvc/routes"

  "github.com/labstack/echo/v4"
)

func main(){
  database.InitDB()

  database.InitialMigration()

  e := echo.New()
  routes.Setup(e)
  e.Logger.Fatal(e.Start(":8000"))
}
