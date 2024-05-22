package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/univearth-rintaro/todo-app/app/routes"
	"github.com/univearth-rintaro/todo-app/db"
)

func main() {
	db.Init()

	server := echo.New()

	// CORS設定を追加
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	routes.RegisterRoutes(server)

	server.Logger.Fatal(server.Start(":5050"))
}
