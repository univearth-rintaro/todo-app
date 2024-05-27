package db

import (
	"fmt"
	"os"

	"github.com/univearth-rintaro/todo-app/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}
}

// InitTestDB initializes the test database
func InitTestDB() *gorm.DB {
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	testDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to test database: " + err.Error())
	}

	// Clean up test database
	testDB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")

	// Run migrations
	err = testDB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("Failed to migrate test database: " + err.Error())
	}

	return testDB
}
