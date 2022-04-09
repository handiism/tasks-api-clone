package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get("user"); !exists {
			response := dtores.NewErrResponse(
				http.StatusBadRequest,
				"Bad Request",
				errors.New("failed to fetch user after jwt middleware"),
			)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}
}
