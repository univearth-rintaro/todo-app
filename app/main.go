package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/univearth-rintaro/todo-app/docs"
	"go.uber.org/zap"

	"github.com/univearth-rintaro/todo-app/api"
	"github.com/univearth-rintaro/todo-app/app/handlers"
	customMiddleware "github.com/univearth-rintaro/todo-app/app/middleware"
	"github.com/univearth-rintaro/todo-app/db"
)

// @title Todo API
// @version 1.0
// @description This is a sample server for a todo application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5050
// @BasePath /

func main() {
	db.Init()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server := echo.New()

	server.Use(customMiddleware.LoggerMiddleware(logger))
	server.Use(echoMiddleware.Recover())

	server.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	apiServer := &handlers.Server{}
	api.RegisterHandlers(server, apiServer)

	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Logger.Fatal(server.Start(":5050"))
}
