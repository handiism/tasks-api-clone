package config

import (
	"fmt"
	"os"

	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`../.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func OpenDBConn() *gorm.DB {
	user := viper.GetString("DB_USER")
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
