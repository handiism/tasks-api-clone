package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handirachmawan/tasks-api-clone/config"
)

func main() {
	db := config.OpenDBConn()
	defer config.CloseDBConn(db)

	r := gin.Default()
	r.Run(":8080")
}
