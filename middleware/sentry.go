package middleware

import (
	"log"
	"net/http"

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
			if p := recover(); p != nil {
				// send error message to sentry
				if hub := sentryecho.GetHubFromContext(c); hub != nil {
					hub.WithScope(func(scope *sentry.Scope) {
						switch x := p.(type) {
						case string:
							hub.CaptureMessage(x)
						case error:
							sentry.CaptureException(x)
						}
					})
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