package main

import (
	"github.com/gin-gonic/gin"
	"tasks/config"
)

func main() {
	db := config.OpenDBConn()
	defer config.CloseDBConn(db)

	r := gin.Default()
	r.Run(":8080")
}
