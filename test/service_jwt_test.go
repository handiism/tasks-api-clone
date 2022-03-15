package test

import (
	"github.com/handirachmawan/tasks-api-clone/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var tokenStr string

func TestServiceJWTGenerate(t *testing.T) {
	jwtService := service.NewJWTService()

	token, err := jwtService.Generate(uuid.New())

	require.Nil(t, err)
	require.NotEmpty(t, token)
	tokenStr = token
}

func TestServiceJWTValidateSuccess(t *testing.T) {
	jwtService := service.NewJWTService()

	token, err := jwtService.Validate(tokenStr, &gin.Context{})

	require.Nil(t, err)
	require.NotEmpty(t, token)
}

func TestServiceJWTValidateFailed(t *testing.T) {
	jwtService := service.NewJWTService()

	token, err := jwtService.Validate("", &gin.Context{})

	require.NotNil(t, err)
	require.Empty(t, token)
}
