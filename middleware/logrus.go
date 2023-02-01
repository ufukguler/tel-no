package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"strings"
)

func LogrusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		uri := request.RequestURI
		if !strings.Contains(uri, "base") && !strings.Contains(uri, "favicon.ico") {

			fields := map[string]interface{}{
				"remoteAddr": request.Header.Get("X-Real-IP"),
				"uri":        request.RequestURI,
				"method":     request.Method,
				"userAgent":  request.Header.Get("User-Agent"),
			}
			logrus.WithFields(fields).Info("[API] ")
		}

		if err := next(c); err != nil {
			logrus.Error(err)
		}
		return nil
	}
}
