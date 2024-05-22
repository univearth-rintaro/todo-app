package db

import (
	"github.com/univearth-rintaro/todo-app/app/models"
)

// Migrate performs the database migrations for all models
func Migrate() {
	err := DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
