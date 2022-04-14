package middleware

import (
	"os"

	"github.com/Investly-id/common-go/v2/payload"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	LogPath *string
	Level   log.Level
}

func NewLogger(LogPath *string, level log.Level) *Logger {
	return &Logger{
		LogPath: LogPath,
		Level:   level,
	}
}

func (l *Logger) makeLogEntry(c echo.Context) *log.Entry {

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
		ForceQuote:      true,
	})

	if l.LogPath != nil {
		f, err := os.OpenFile(*l.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Panic(err)
		}

		// set output to file and make log to json
		log.SetOutput(f)
		log.SetFormatter(&log.JSONFormatter{})
	}

	if c == nil {
		return log.WithFields(log.Fields{})
	}

	return log.WithFields(log.Fields{
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func (l *Logger) MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l.makeLogEntry(c).Info("Incoming request")
		return next(c)
	}
}

func (l *Logger) ErrorHandler(err error, c echo.Context) {

	report := err.(*echo.HTTPError)

	resp := &payload.Response{
		Message: report.Message.(string),
		Error:   report.Message.(string),
	}

	l.makeLogEntry(c).Error(report.Message)
	c.JSON(report.Code, &resp)
}
