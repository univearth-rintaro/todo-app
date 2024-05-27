package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/univearth-rintaro/todo-app/api"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetTodos(ctx echo.Context) error {
	todos := []api.Todo{}
	if err := s.DB.Find(&todos).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todos)
}

func (s *Server) PostTodos(ctx echo.Context) error {
	var todo api.Todo
	if err := ctx.Bind(&todo); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	if err := s.DB.Create(&todo).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, todo)
}

func (s *Server) GetTodosId(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := s.DB.First(&todo, id).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (s *Server) PutTodosId(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := ctx.Bind(&todo); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	if err := s.DB.Model(&api.Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (s *Server) DeleteTodosId(ctx echo.Context, id int) error {
	if err := s.DB.Delete(&api.Todo{}, id).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) PatchTodosIdDone(ctx echo.Context, id int) error {
	var todo api.Todo
	if err := s.DB.First(&todo, id).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	done := true
	todo.Done = &done
	if err := s.DB.Save(&todo).Error; err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

type ErrorResponse struct {
	Error   string `json:"error"`
	TraceID string `json:"trace_id"`
}

func handleError(ctx echo.Context, statusCode int, err error) error {
	traceID := ctx.Get("traceID").(string)
	response := ErrorResponse{
		Error:   err.Error(),
		TraceID: traceID,
	}
	return ctx.JSON(statusCode, response)
}
