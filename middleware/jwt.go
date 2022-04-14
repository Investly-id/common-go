package middleware

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtMiddleware struct {
	AccessTokenExpiredHour int
	Secret                 string
}

func NewJwt(expiredHour int, secret string) *JwtMiddleware {
	return &JwtMiddleware{
		AccessTokenExpiredHour: expiredHour,
		Secret:                 secret,
	}
}

type jwtCustomClaims struct {
	jwt.StandardClaims
}

func (j *JwtMiddleware) GenerateToken(userId int) (string, error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			Subject:   strconv.Itoa(userId),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.AccessTokenExpiredHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.Secret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *JwtMiddleware) ValidateJWT() echo.MiddlewareFunc {

	JWTConfig := middleware.JWTConfig{
		TokenLookup: "header:" + echo.HeaderAuthorization,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return []byte(j.Secret), nil
			}

			token, err := jwt.Parse(auth, keyFunc)

			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	return middleware.JWTWithConfig(JWTConfig)
}

func (j *JwtMiddleware) GetJWTClaims(c echo.Context) map[string]interface{} {
	jwtContext := c.Get("user").(*jwt.Token)
	return jwtContext.Claims.(jwt.MapClaims)
}

func (j *JwtMiddleware) GetUserIdFromJwt(c echo.Context) int {
	userId, _ := strconv.Atoi(j.GetJWTClaims(c)["sub"].(string))
	return userId
}
