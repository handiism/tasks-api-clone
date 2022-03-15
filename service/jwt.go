package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTService interface {
	Generate(uuid uuid.UUID) (string, error)
	Validate(token string, ctx *gin.Context) (*jwt.Token, error)
}

type customClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "Handi Rachmawan",
		secretKey: "handiism",
	}
}

func (j *jwtService) Generate(uuid uuid.UUID) (string, error) {
	claims := &customClaims{
		UUID: uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issuer,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(24 * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func (j *jwtService) Validate(token string, ctx *gin.Context) (*jwt.Token, error) {
	custClaims := &customClaims{}

	t, err := jwt.ParseWithClaims(token, custClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*customClaims); ok && t.Valid {
		if err := claims.Valid(); err != nil {
			return nil, err
		}

		if ok := claims.VerifyIssuer("Handi Rachmawan", true); !ok {
			return nil, errors.New("invaid jwt issuer")
		}

		if ok := claims.VerifyExpiresAt(time.Now(), true); !ok {
			return nil, errors.New("token is expired")
		}
	} else if ok && !t.Valid {
		return nil, errors.New("invalid token")
	} else if !ok {
		return nil, errors.New("unexpected jwt claims format")
	}

	return t, nil
}
