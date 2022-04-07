package config

import (
	"fmt"

	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConn() *gorm.DB {
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	hostname := viper.GetString("DB_HOSTNAME")
	port := viper.GetString("DB_PORT")

	database := viper.GetString("DB_DATABASE")

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
