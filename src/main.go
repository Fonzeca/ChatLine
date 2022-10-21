package main

import (
	"github.com/Fonzeca/Chatline/src/entry"
	"github.com/Fonzeca/Chatline/src/server/api"
	"github.com/Fonzeca/Chatline/src/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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

	e.Logger.Fatal(e.Start(":6548"))
}
