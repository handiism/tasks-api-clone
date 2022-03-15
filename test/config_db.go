package test

import (
	"fmt"
	"github.com/handirachmawan/tasks-api-clone/helper"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openDBConn() *gorm.DB {
	helper.GetEnv()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")

	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		hostname, user, password, port, database+"_test",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func closeDBConn(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}
	conn.Close()
}
