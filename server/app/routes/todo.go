package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/univearth-rintaro/todo-app/app/models"
	"github.com/univearth-rintaro/todo-app/app/services"
)

func GetAllTodos(context echo.Context) error {
	todos, err := services.GetAllTodos()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, todos)
}

func GetTodoById(context echo.Context) error {
	todoId := context.Param("todoId")
	todo, err := services.GetTodoById(todoId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, todo)
}

func CreateTodo(context echo.Context) error {
	todo := new(models.Todo)
	if err := context.Bind(todo); err != nil {
		return context.JSON(http.StatusBadRequest, "リクエストの内容が正しくありません。")
	}
	if err := services.CreateTodo(todo); err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを作成できませんでした。")
	}
	return context.JSON(http.StatusCreated, todo)
}

func DeleteTodoById(context echo.Context) error {
	todoId := context.Param("todoId")
	if err := services.DeleteTodoById(todoId); err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを削除できませんでした。")
	}
	return context.NoContent(http.StatusNoContent)
}

func UpdateTodoById(context echo.Context) error {
	todoId := context.Param("todoId")
	updatedTodo := new(models.Todo)
	if err := context.Bind(updatedTodo); err != nil {
		return context.JSON(http.StatusBadRequest, "リクエストの内容が正しくありません。")
	}
	if err := services.UpdateTodoById(todoId, updatedTodo); err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを更新できませんでした。")
	}
	return context.JSON(http.StatusOK, updatedTodo)
}

func MarkTodoAsDone(context echo.Context) error {
	todoId := context.Param("todoId")
	if err := services.MarkTodoAsDone(todoId); err != nil {
		return context.JSON(http.StatusInternalServerError, "todoを完了に設定できませんでした。")
	}
	return context.NoContent(http.StatusNoContent)
}
