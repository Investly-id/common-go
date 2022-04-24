package middleware

import (
	"net/http"

	"github.com/Investly-id/common-go/v3/payload"
	"github.com/labstack/echo/v4"
)

type InternalConnection struct {
	secret string
}

func NewInternalConnection(secret string) *InternalConnection {
	return &InternalConnection{
		secret: secret,
	}
}

func (icm *InternalConnection) ValidateInternalAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// get inernal secret from headers
		internalToken := c.Request().Header.Get("x-internal-token")

		// check if request doesn't have x-internal-token headers
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
