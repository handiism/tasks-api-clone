package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	dtores "github.com/handiism/tasks-api-clone/dto/response"
	"github.com/handiism/tasks-api-clone/repo"
	"github.com/handiism/tasks-api-clone/service"
)

func AuthorizeJWT(userRepo repo.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := dtores.NewErrReponse(
				http.StatusUnauthorized,
				"Unauthorized",
				errors.New("request header does'nt contain bearer token"),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenStr string
		arr := strings.Split(authHeader, " ")
		jwtService := service.NewJWTService()

		if len(arr) == 2 {
			tokenStr = arr[1]
		}

		token, err := jwtService.Validate(tokenStr)

		if err != nil {
			response := dtores.NewErrReponse(http.StatusUnauthorized, "Unauthorized", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := token.Claims.(*service.UUIDClaims)

		if !ok || !token.Valid {
			response := dtores.NewErrReponse(http.StatusUnauthorized, "Unauthorized", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		uuid, err := uuid.Parse(claims.UUID)

		if err != nil {
			response := dtores.NewErrReponse(http.StatusUnauthorized, "Unauthorized", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userRepo.FindByUUID(uuid)
		if err != nil {
			response := dtores.NewErrReponse(http.StatusUnauthorized, "Unauthorized", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, _ = userRepo.PreloadAll(user)

		c.Set("user", user)
	}
}
