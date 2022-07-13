// Package logger is a package to simplify logging in Go.
package logger

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	log.FieldLogger
}

type logger struct {
	log.FieldLogger
}

func NewLogger() Logger {
	logg := log.New()

	if os.Getenv("ENVIRONMENT") == "production" {
		logg.SetFormatter(&log.JSONFormatter{})
	} else {
		logg.SetFormatter(&log.TextFormatter{})
	}

	entry := log.NewEntry(logg)
	entry = entry.WithFields(log.Fields{
		"index": "effective-eureka",
	})

	return &logger{
		entry,
	}
}

// Middleware provides a echo.MiddlewareFunc to use in echo server with logrus.
func Middleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// set up new log entry
			logg := NewLogger()

			// set up log event duration field
			logEntry := logg.WithFields(log.Fields{
				"event": log.Fields{
					"duration": time.Since(v.StartTime),
				},
			})

			// add http request fields in production environment
			// in development can be ignored to clean logs
			if os.Getenv("ENVIRONMENT") == "production" {
				logEntry = logEntry.WithFields(log.Fields{
					"http": log.Fields{
						"request": log.Fields{
							"method":     v.Method,
							"path":       v.URI,
							"status":     v.Status,
							"latency":    v.Latency,
							"user-agent": v.UserAgent,
							"remote-ip":  v.RemoteIP,
						},
					},
				})
			}

			logEntry.Info("request handled")

			return nil
		},
	})
}
