package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"telno/config"
	"telno/models"
)

var e models.Response

func AdminCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.String(), "/api/admin") && c.Request().Method != http.MethodOptions {
			if c.Request().Header.Get("admin") != config.ApiConfig.AdminPass { // TODO FIXME
				return e.ResponseUnauthorized(c)
			}

		}
		return next(c)
	}
}
