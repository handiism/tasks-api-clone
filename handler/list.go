package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/service"
)

type listHandler struct {
	listService service.ListService
	userService service.UserService
}

func NewListHandler(listService service.ListService, userService service.UserService) listHandler {
	return listHandler{
		listService: listService,
		userService: userService,
	}
}

func (l *listHandler) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp dtoreq.SaveList
		err := c.ShouldBindJSON(&inp)
		if err != nil {
			resp := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		user, _ := c.Get("user")
		if list, err := l.listService.Add(user.(model.User).ID, inp); err != nil {
			resp := dtores.NewErrResponse(
				http.StatusUnprocessableEntity,
				"Unprocessable Entitiy",
				err,
			)
			c.JSON(http.StatusUnprocessableEntity, resp)
		} else {
			resp := dtores.NewResponse(http.StatusCreated, "Created", list)
			c.JSON(http.StatusCreated, resp)
		}
	}
}

func (l *listHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		listIdStr := c.Param("id")
		listId, err := strconv.Atoi(listIdStr)

		if err != nil {
			resp := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		list, err := l.listService.Fetch(user.(model.User).ID, uint(listId))
		if err != nil {
			resp := dtores.NewErrResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err)
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		resp := dtores.NewResponse(http.StatusOK, "OK", list)
		c.JSON(http.StatusOK, resp)
	}
}

func (l *listHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp dtoreq.SaveList
		err := c.ShouldBindJSON(&inp)
		if err != nil {
			resp := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		user, _ := c.Get("user")
		listIdSTr := c.Param("id")
		listId, err := strconv.Atoi(listIdSTr)
		if err != nil {
			resp := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		if res, err := l.listService.Update(user.(model.User).ID, uint(listId), inp); err != nil {
			resp := dtores.NewErrResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err)
			c.JSON(http.StatusUnprocessableEntity, resp)
		} else {
			resp := dtores.NewResponse(http.StatusOK, "OK", res)
			c.JSON(http.StatusOK, resp)
		}
	}
}

func (l *listHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		if lists, err := l.listService.FetchAll(user.(model.User).ID); err != nil {
			resp := dtores.NewErrResponse(http.StatusUnprocessableEntity, "Unprocessable Entitiy", err)
			c.JSON(http.StatusUnprocessableEntity, resp)
		} else {
			resp := dtores.NewResponse(http.StatusOK, "OK", lists)
			c.JSON(http.StatusOK, resp)
		}

	}
}

func (l *listHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		listIdStr := c.Param("id")
		if listId, err := strconv.Atoi(listIdStr); err != nil {
			resp := dtores.NewErrResponse(http.StatusBadRequest, "Bad Request", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		} else {
			if err := l.listService.Delete(user.(model.User).ID, uint(listId)); err != nil {
				resp := dtores.NewErrResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err)
				c.JSON(http.StatusUnprocessableEntity, resp)
			} else {
				resp := dtores.NewResponse(http.StatusOK, "OK", nil)
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}
