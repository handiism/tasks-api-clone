package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/handiism/tasks-api-clone/config"
	"github.com/handiism/tasks-api-clone/handler"
	"github.com/handiism/tasks-api-clone/middleware"
	"github.com/handiism/tasks-api-clone/repo"
	"github.com/handiism/tasks-api-clone/service"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	db := config.OpenDBConn()
	defer config.CloseDBConn(db)

	userRepo := repo.NewUserRepo(db)
	jwtHandler := middleware.AuthorizeJWT(userRepo)
	userService := service.NewUserService(userRepo)
	jwtService := service.NewJWTService()
	userHandler := handler.NewUserHandler(
		userService,
		jwtService,
	)

	r := gin.Default()

	r.POST("register", userHandler.Register())
	r.POST("login", userHandler.Login())

	userRoutes := r.Group("user", jwtHandler)
	{
		userRoutes.GET("/", userHandler.Fetch())
		userRoutes.PATCH("/", userHandler.Update())
		userRoutes.DELETE("/", userHandler.Delete())
	}

	r.Run(fmt.Sprintf("%s:%s",
		viper.GetString("SERVER_HOST"),
		viper.GetString("SERVER_PORT"),
	))
}
