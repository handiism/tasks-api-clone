package main

import (
	"google-tasks-clone/config"
	"google-tasks-clone/helper"
	"google-tasks-clone/model"
)

func main() {

	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	err := db.AutoMigrate(
		&model.User{},
		&model.List{},
		&model.Task{},
		&model.Subtask{},
	)

	helper.DefaultPanic(err)
}
