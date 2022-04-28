package middleware

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Hook interface {
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
}

type hook struct {
	ctx echo.Context
}

func (h *hook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (h *hook) Fire(entry *logrus.Entry) error {
	if entry.Level == logrus.ErrorLevel {
		if hub := sentryecho.GetHubFromContext(h.ctx); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetExtra("TEST ERROR EXTRA", "TEST ERROR EXTRA")
				hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
			})
		}
	}

	defer sentry.Flush(2 * time.Second)

	sentryID := sentry.CaptureMessage("Captue message yes")
	log.Println("sentry ID", sentryID)

	return nil
}

type Sentry struct {
	Dsn         string
	Mid         echo.MiddlewareFunc
	LogrusHooks Hook
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
		return err
	}

	s.Dsn = dsn
	return nil
}

func (s *Sentry) AddLogrusHook(ctx echo.Context) {
	s.LogrusHooks = &hook{
		ctx: ctx,
	}
}
