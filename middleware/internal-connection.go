package middleware

import (
	"net/http"

	"github.com/Investly-id/common-go/v2/payload"
	"github.com/labstack/echo/v4"
)

type InternalConnectionMiddleware struct {
	secret string
}

func NewInternalConnection(secret string) *InternalConnectionMiddleware {
	return &InternalConnectionMiddleware{
		secret: secret,
	}
}

func (icm *InternalConnectionMiddleware) ValidateInternalAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		internalToken := c.Request().Header.Get("x-internal-token")
		if internalToken != icm.secret {
			res := &payload.Response{
				Message: "Unauthorized",
				Status:  false,
			}
			return c.JSON(http.StatusUnauthorized, res)
		}
		return next(c)
	}
}
