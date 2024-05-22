package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.GET("/todos", GetAllTodos)
	server.GET("/todos/:todoId", GetTodoById)
	server.POST("/todos", CreateTodo)
	server.DELETE("/todos/:todoId", DeleteTodoById)
	server.PATCH("/todos/:todoId", UpdateTodoById)
	server.PATCH("/todos/:todoId/done", MarkTodoAsDone)
}
