package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/univearth-rintaro/todo-app/api"
	"github.com/univearth-rintaro/todo-app/db"
)

type Server struct{}

func (s *Server) GetTodos(ctx echo.Context) error {
	todos := []api.Todo{}
	if err := db.DB.Find(&todos).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.JSON(200, todos)
}

func (s *Server) PostTodos(ctx echo.Context) error {
	var todo api.Todo
	if err := ctx.Bind(&todo); err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	if err := db.DB.Create(&todo).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.JSON(201, todo)
}

func (s *Server) GetTodosId(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.JSON(200, todo)
}

func (s *Server) PutTodosId(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := ctx.Bind(&todo); err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	if err := db.DB.Model(&api.Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.JSON(200, todo)
}

func (s *Server) DeleteTodosId(ctx echo.Context, id int) error {
	if err := db.DB.Delete(&api.Todo{}, id).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.NoContent(204)
}

func (s *Server) PatchTodosIdDone(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	done := true
	todo.Done = &done
	if err := db.DB.Save(&todo).Error; err != nil {
		traceID := ctx.Get("traceID").(string)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error":    err.Error(),
			"trace_id": traceID,
		})
	}
	return ctx.JSON(200, todo)
}
