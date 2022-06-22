package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/model"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, exist := c.Get("user"); !exist {
			response := dtores.NewErrResponse(
				http.StatusBadRequest,
				"Bad Request",
				errors.New("failed to fetch user after jwt middleware"),
			)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			if user, ok := val.(model.User); ok {
				list := user.Lists
				c.Set("lists", list)
			} else {
				response := dtores.NewErrResponse(
					http.StatusBadRequest,
					"Bad Request",
					errors.New("unrecognized user context"),
				)
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			}
		}
	}
}
