package main

import (
	"fmt"

	"github.com/Fonzeca/Chatline/src/entry"
	"github.com/Fonzeca/Chatline/src/server/api"
	"github.com/Fonzeca/Chatline/src/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	// go monitor.System()
	_, closeFunc := services.SetupRabbitMq()
	defer closeFunc()
	entry.NewRabbitMqDataEntry()

	e := echo.New()

	api := api.NewApiComment()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.POST("/createComment", api.CreateComment)
	e.GET("/getAllComments", api.GetAllComments)
	e.POST("/getCommentsByUserIds", api.GetCommentsByUserIds)

	e.Logger.Fatal(e.Start(":4762"))
}

func InitConfig() {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("\nfatal error config file: %w", err))
	}
}
