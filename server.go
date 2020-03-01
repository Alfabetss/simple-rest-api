package main

import (
	"fmt"
	"log"

	"github.com/Alfabetss/simple-rest-api/config"
	"github.com/Alfabetss/simple-rest-api/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Printf("failed to read config file : %s", err.Error())
		log.Fatal()
	}

	db, err := config.Connect(cfg)
	if err != nil {
		log.Printf("failed to connect database : %s", err.Error())
		log.Fatal()
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	talentController := controller.NewTalentController(cfg, db)
	talentGroup := e.Group("/talent")
	talentGroup.POST("", talentController.CreateTalent)
	talentGroup.GET("/:talentID", talentController.FindTalent)
	talentGroup.DELETE("/:talentID", talentController.Delete)
	talentGroup.PUT("", talentController.UpdateTalent)
	talentGroup.PUT("/experience", talentController.UpdateTalentExperience)

	if err := e.Start(cfg.Server.Port); err != nil {
		log.Fatal()
	}
}

func readConfig() (config.Configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var config config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config")
	}

	// set default port
	viper.SetDefault("server.Port", ":8888")
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to bind config")
	}

	fmt.Println("success read from config")
	fmt.Printf("Database name : %s", config.Database.DBName)
	return config, nil
}
