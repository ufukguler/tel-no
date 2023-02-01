package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"telno/config"
)

// GetSitemapXML main sitemap file
func GetSitemapXML(c echo.Context) error {
	return c.Blob(http.StatusOK, "text/xml; charset=utf-8", config.ApiConfig.SitemapMain)
}

// GetSitemapPages sitemap top pages file
func GetSitemapPages(c echo.Context) error {
	return c.Blob(http.StatusOK, "text/xml; charset=utf-8", config.ApiConfig.Sitemap.XMLContent())
}

// GetSitemap sitemap file that contains phone numbers
func GetSitemap(c echo.Context) error {
	param := strings.ReplaceAll(c.Param("id"), ".xml", "")
	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil || id < 1 || id > int64(len(config.ApiConfig.Sitemaps)) {
		return e.ResponseError(errors.New("invalid number"), c)
	}

	sitemap := config.ApiConfig.Sitemaps[id-1]
	return c.Blob(http.StatusOK, "text/xml; charset=utf-8", sitemap.XMLContent())
}
