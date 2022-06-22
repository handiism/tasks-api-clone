package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/service"
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
		var input dtoreq.SaveUser
		err := c.ShouldBindJSON(&input)
		if err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnprocessableEntity,
				"Unprocessable Entity",
				err,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		user, err := u.userService.Register(input)
		if err != nil {
			response := dtores.NewErrResponse(
				http.StatusBadRequest, "Bad Request", err,
			)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		user.Token, err = u.jwtService.Generate(user.ID)
		if err != nil {
			response := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := dtores.NewResponse(
			http.StatusCreated, "Created", user,
		)
		c.JSON(http.StatusCreated, response)
	}
}

func (u *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp dtoreq.Login
		err := c.ShouldBindJSON(&inp)
		if err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnprocessableEntity, "Unprocessable Entity", err,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		user, err := u.userService.Login(inp)
		if err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnauthorized, "Unauthorized", err,
			)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		user.Token, err = u.jwtService.Generate(user.ID)
		if err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnauthorized, "Unauthorized", err,
			)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		response := dtores.NewResponse(
			http.StatusOK, "OK", user,
		)
		c.JSON(http.StatusOK, response)
	}
}

func (u *userHandler) Fetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")
		if !exist {
			response := dtores.NewErrResponse(
				http.StatusNotFound,
				"Not Found",
				errors.New("user can't be processed"),
			)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := dtores.NewResponse(
			http.StatusOK, "OK", user,
		)
		c.JSON(http.StatusOK, response)
	}
}

func (u *userHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp dtoreq.SaveUser
		if err := c.ShouldBind(&inp); err != nil {
			response := dtores.NewResponse(
				http.StatusBadRequest, "Bad Request", err,
			)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		userMap, exist := c.Get("user")
		if !exist {
			response := dtores.NewErrResponse(
				http.StatusNotFound,
				"Not Found",
				errors.New("user not found"),
			)
			c.JSON(http.StatusNotFound, response)
			return
		}
		user, ok := userMap.(model.User)
		if !ok {
			response := dtores.NewResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				errors.New("failed to convert map[string]interface{} to model.User{}"),
			)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		if user, err := u.userService.Update(user.ID, inp); err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnprocessableEntity, "Unprocessable Entity", err,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
		} else {
			response := dtores.NewResponse(http.StatusOK, "OK", user)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (u *userHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		userAny, exist := c.Get("user")
		if !exist {
			response := dtores.NewErrResponse(
				http.StatusNotFound,
				"Not Found",
				errors.New("key user not found in middleware context"),
			)
			c.JSON(http.StatusNotFound, response)
			return
		}

		user, ok := userAny.(model.User)
		if !ok {
			response := dtores.NewErrResponse(
				http.StatusUnprocessableEntity, "Unprocessable Entity",
				errors.New("cannot cast key user to model user"),
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		if err := u.userService.Delete(user.ID.String()); err != nil {
			response := dtores.NewErrResponse(
				http.StatusUnprocessableEntity, "Unprocessable Entity", err,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
		} else {
			response := dtores.NewResponse(
				http.StatusOK, "OK", nil,
			)
			c.JSON(http.StatusOK, response)
		}
	}
}
