package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			stop := time.Now()

			fields := []zap.Field{
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.String("remote_ip", c.RealIP()),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", stop.Sub(start)),
			}

			if err != nil {
				fields = append(fields, zap.Error(err))
				c.Error(err)
			}

			logger.Info("access log", fields...)

			return nil
		}
	}
}
