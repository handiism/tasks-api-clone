package repo_test

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	viper.SetConfigFile(`../.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func openDBConn() *gorm.DB {
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	hostname := viper.GetString("DB_HOSTNAME")
	port := viper.GetString("DB_PORT")

	database := viper.GetString("DB_DATABASE_TEST")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		hostname, user, password, port, database,
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
