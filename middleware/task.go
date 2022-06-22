package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/model"
)

func AuthorizeList() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, exist := c.Get("lists"); !exist {
			response := dtores.NewErrResponse(
				http.StatusBadRequest,
				"Bad Request",
				errors.New("failed to fetch all lists after list middleware"),
			)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			if lists, ok := val.([]model.List); ok {
				listIdString := c.Param("list_id")
				listId, err := strconv.Atoi(listIdString)

				if err != nil {
					response := dtores.NewErrResponse(
						http.StatusBadRequest,
						"Bad Request",
						errors.New("failed to fetch list with id "+listIdString),
					)
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				}

				found := false
				index := 0
				for k, v := range lists {
					if int(v.ID) == listId {
						found = true
						index = k
						break
					} else {
						continue
					}
				}
				if !found {
					response := dtores.NewErrResponse(
						http.StatusNotFound,
						"Not Found",
						errors.New("list not found in targeted user"),
					)
					c.AbortWithStatusJSON(http.StatusNotFound, response)
				} else {
					c.Set("tasks", lists[index].Tasks)
				}
			}
		}
	}
}
