// Package logger is a package to simplify logging in Go.
package logger

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

// Logger is a interface who provides methods to manage logs.
type Logger interface {
	// log.FieldLogger composes Logger with log.FieldLogger from logrus
	log.FieldLogger
}

// logger is a struct who implements Logger interface.
type logger struct {
	log.FieldLogger
}

// NewLogger returns a new instance of Logger.
func NewLogger() Logger {
	// set up new logger from logrus
	logg := log.New()

	// set formatter as JSONFormatter for production environment
	// while in another env use TextFormatter to more clean logs to read
	// and debug
	if os.Getenv("ENVIRONMENT") == "production" {
		logg.SetFormatter(&log.JSONFormatter{})
	} else {
		logg.SetFormatter(&log.TextFormatter{})
	}

	// set up log entry from logg instance
	entry := log.NewEntry(logg)
	// initiate index to easy send in elastic opensearch index
	// with log sender lambda in AWS
	entry = entry.WithFields(log.Fields{
		"index": "effective-eureka",
	})

	// return new instance of logger with inner entry
	return &logger{
		// inner entry provides a better way to log with additional
		// fields from logrus
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
