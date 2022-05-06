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
	Mid echo.MiddlewareFunc
}

func NewSentry() *Sentry {
	return &Sentry{
		Mid: sentryecho.New(sentryecho.Options{
			Repanic: true,
		}),
	}
}

func (s *Sentry) InitSentryConnection(dsn string) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	}); err != nil {
		log.Fatal(err)
		return err
	}

	s.Dsn = dsn
	return nil
}

func (s *Sentry) Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			// recover from panic
			if err := recover(); err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
				c.JSON(http.StatusInternalServerError, &payload.Response{
					Status:  false,
					Message: "Internal server error",
				})
			}
		}()
		return next(c)
	}
}
