package config

import (
	"fmt"
	"github.com/handirachmawan/tasks-api-clone/helper"
	"github.com/handirachmawan/tasks-api-clone/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConn() *gorm.DB {
	helper.GetEnv()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")

	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		hostname, user, password, port, database,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&model.User{}, &model.List{}, &model.Task{}, &model.Subtask{},
	); err != nil {
		panic(err)
	}

	return db
}

func CloseDBConn(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	conn.Close()
}
