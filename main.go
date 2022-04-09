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
	listRepo := repo.NewListRepo(db)

	jwtMiddleware := middleware.AuthorizeJWT(userRepo)
	listMiddleware := middleware.AuthorizeUser()

	userService := service.NewUserService(userRepo)
	listService := service.NewListService(userRepo, listRepo)
	jwtService := service.NewJWTService()

	userHandler := handler.NewUserHandler(
		userService,
		jwtService,
	)
	listHandler := handler.NewListHandler(
		listService,
		userService,
	)

	r := gin.Default()

	r.POST("register", userHandler.Register())
	r.POST("login", userHandler.Login())

	userRoutes := r.Group("user", jwtMiddleware)
	{
		userRoutes.GET("/", userHandler.Fetch())
		userRoutes.PATCH("/", userHandler.Update())
		userRoutes.DELETE("/", userHandler.Delete())
	}

	listRoutes := r.Group("user/list", jwtMiddleware, listMiddleware)
	{
		listRoutes.POST("/", listHandler.Add())
		listRoutes.GET("/:id", listHandler.Get())
	}

	r.Run(fmt.Sprintf("%s:%s",
		viper.GetString("SERVER_HOST"),
		viper.GetString("SERVER_PORT"),
	))
}
