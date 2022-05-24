package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/Investly-id/common-go/v3/payload"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

type Sentry struct {
	Dsn string
	Env string
	Mid echo.MiddlewareFunc
}

func NewSentry(dsn string, env string) *Sentry {

	// init sentry when env is production
	if env == "prod" || env == "production" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         dsn,
			Environment: env,
		}); err != nil {
			log.Fatal(err)
			return nil
		}
	}

	return &Sentry{
		Mid: sentryecho.New(sentryecho.Options{}),
		Dsn: dsn,
		Env: env,
	}
}

func (s *Sentry) Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			// recover from panic
			if err := recover(); err != nil {
				// push sentry error message
				if s.Env == "prod" || s.Env == "production" {
					sentry.CurrentHub().Recover(err)
					sentry.Flush(time.Second * 5)
				}

				c.JSON(http.StatusInternalServerError, &payload.Response{
					Status:  false,
					Message: "Internal server error",
				})
			}
		}()
		return next(c)
	}
}
