package config

import (
	"fmt"
	"google-tasks-clone/helper"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	helper.CustomPanic(errEnv, "Failed to load .env file")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		hostname,
		port,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.CustomPanic(err, "Failed to create connection to database")
	// db.AutoMigrate()
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.CustomPanic(err, "Failed to close database connection")
	dbSQL.Close()
}
