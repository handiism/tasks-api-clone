package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtoreq "github.com/handirachmawan/tasks-api-clone/dto/request"
	dtores "github.com/handirachmawan/tasks-api-clone/dto/response"
	"github.com/handirachmawan/tasks-api-clone/service"
)

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserHandler(
	userService service.UserService,
	jwtService service.JWTService,
) *userHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (u *userHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dtoreq.Register
		err := c.ShouldBindJSON(&input)
		if err != nil {
			errors := dtores.ValidationErrors(err)
			data := gin.H{"errors": errors}
			response := dtores.NewResponse(
				"Failed to register an account",
				http.StatusUnprocessableEntity,
				"error",
				data,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		user, err := u.userService.Register(input)
		if err != nil {
			response := dtores.NewResponse(
				"Failed to register an account",
				http.StatusBadRequest,
				"error",
				nil,
			)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		user.Token, err = u.jwtService.Generate(user.ID)
		if err != nil {
			response := dtores.NewResponse(
				"Failed to register an account",
				http.StatusBadRequest,
				"error",
				nil,
			)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := dtores.NewResponse(
			"Account has been registered",
			http.StatusOK,
			"success",
			user,
		)
		c.JSON(http.StatusOK, response)
	}
}

func (u *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp dtoreq.Login
		err := c.ShouldBindJSON(&inp)
		if err != nil {
			errors := dtores.ValidationErrors(err)
			data := gin.H{"errors": errors}
			response := dtores.NewResponse(
				"Failed to register an account",
				http.StatusUnprocessableEntity,
				"error",
				data,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		user, err := u.userService.Login(inp)
		if err != nil {
			data := gin.H{"errors": err}
			response := dtores.NewResponse(
				"Login Failed",
				http.StatusUnauthorized,
				"errors",
				data,
			)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		user.Token, err = u.jwtService.Generate(user.ID)
		if err != nil {
			data := gin.H{"errors": err}
			response := dtores.NewResponse(
				"Login Failed",
				http.StatusUnauthorized,
				"errors",
				data,
			)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		response := dtores.NewResponse(
			"Login Success",
			http.StatusOK,
			"success",
			user,
		)
		c.JSON(http.StatusOK, response)
	}
}

func (u *userHandler) Fetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")
		if !exist {
			response := dtores.NewResponse("User did not fetched", http.StatusNotFound, "error", nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := dtores.NewResponse("User fetched", http.StatusOK, "success", user)
		c.JSON(http.StatusOK, response)
	}
}
