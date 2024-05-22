package services

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/univearth-rintaro/todo-app/app/models"
	"github.com/univearth-rintaro/todo-app/db"
)

func GetAllTodos() ([]models.Todo, error) {
	todos := []models.Todo{}
	if err := db.DB.Find(&todos).Error; err != nil {
		log.Println("Error finding todos:", err) // ここでエラーログを追加
		return nil, echo.ErrNotFound
	}
	return todos, nil
}

func GetTodoById(id string) (*models.Todo, error) {
	todo := models.Todo{}
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		log.Println("Error finding todo by ID:", err)
		return nil, echo.ErrNotFound
	}
	return &todo, nil
}

func CreateTodo(todo *models.Todo) error {
	if err := db.DB.Create(&todo).Error; err != nil {
		log.Println("Error creating todo:", err)
		return err
	}
	return nil
}

func DeleteTodoById(id string) error {
	if err := db.DB.Delete(&models.Todo{}, id).Error; err != nil {
		log.Println("Error deleting todo:", err)
		return err
	}
	return nil
}

func UpdateTodoById(id string, updatedTodo *models.Todo) error {
	todo := models.Todo{}
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		log.Println("Error finding todo for update:", err)
		return echo.ErrNotFound
	}
	if err := db.DB.Model(&todo).Updates(updatedTodo).Error; err != nil {
		log.Println("Error updating todo:", err)
		return err
	}
	return nil
}

func MarkTodoAsDone(id string) error {
	todo := models.Todo{}
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		log.Println("Error finding todo for marking as done:", err)
		return echo.ErrNotFound
	}
	todo.Done = true
	if err := db.DB.Save(&todo).Error; err != nil {
		log.Println("Error marking todo as done:", err)
		return err
	}
	return nil
}
