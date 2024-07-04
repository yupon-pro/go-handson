package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yupon-pro/go-handson/work/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.PostUser)
	e.PATCH("/users/:id", handlers.PatchUser)
	e.DELETE("/users/:id", handlers.DeleteUser)


	e.Logger.Fatal(e.Start(":8080"))
}