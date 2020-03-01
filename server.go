package main

import (
	"log"

	"github.com/Alfabetss/simple-rest-api/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	talentGroup := e.Group("/talent")
	talentGroup.POST("/create", controller.CreateTalent)
	talentGroup.GET("/:talentID", controller.FindTalent)
	talentGroup.DELETE("/:talentID", controller.Delete)

	if err := e.Start(":1122"); err != nil {
		log.Fatal()
	}
}
