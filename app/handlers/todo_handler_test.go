package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/univearth-rintaro/todo-app/app/models"
	"github.com/univearth-rintaro/todo-app/db"
)

func TestGetTodos(t *testing.T) {
	// Set up Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up test database
	testDB := db.InitTestDB()
	defer func() {
		sqlDB, _ := testDB.DB()
		sqlDB.Close()
	}()

	// Insert a test todo
	testTodo := &models.Todo{
		Title:       "Test Todo",
		Description: "This is a test todo.",
		Status:      "pending",
	}
	testDB.Create(testTodo)

	// Instantiate handler with test database
	h := &Server{DB: testDB}

	// Assertions
	if assert.NoError(t, h.GetTodos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Todo")
	}
}
