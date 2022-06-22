package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/model"
)

func AuthorizeTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, exist := c.Get("tasks"); !exist {
			response := dtores.NewErrResponse(
				http.StatusBadRequest,
				"Bad Request",
				errors.New("failed to fetch all tasks after tasks middleware"),
			)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			_, ok := val.([]model.Task)

			if !ok {
				// Handle an error
			}

			// Get a param task_id
			// If param not found, handle an error
		}
	}
}
