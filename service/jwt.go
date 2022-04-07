package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type JWTService interface {
	Generate(uuid uuid.UUID) (string, error)
	Validate(token string) (*jwt.Token, error)
}

type UUIDClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    viper.GetString("JWT_ISSUER"),
		secretKey: viper.GetString("JWT_SECRET_KEY"),
	}
}

func (j *jwtService) Generate(id uuid.UUID) (string, error) {
	var nullId uuid.UUID
	if id == nullId {
		return "", errors.New("null uuid")
	}

	claims := &UUIDClaims{
		UUID: id.String(),
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

func (j *jwtService) Validate(token string) (*jwt.Token, error) {
	custClaims := &UUIDClaims{}

	t, err := jwt.ParseWithClaims(token, custClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*UUIDClaims); ok && t.Valid {
		if err := claims.Valid(); err != nil {
			return nil, err
		}

		if ok := claims.VerifyIssuer(j.issuer, true); !ok {
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
