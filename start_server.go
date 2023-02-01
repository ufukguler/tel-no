package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"telno/config"
	"telno/database"
	m "telno/middleware"
	"telno/models"
	routes "telno/route"
	"telno/service"
	"telno/service/latest_bar"
	"telno/service/sitemap"
	"time"
)

func initServer(e *echo.Echo) {

	config.LoadEnv(".env")

	config.LoadConfig()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(m.LogrusMiddleware)
	e.Use(m.AfterResponseMiddleware)
	e.Use(m.AdminCheck)
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace},
	}))
	e.Any("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Not Found",
			Data:    nil,
		})
	})
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	// LOGGER
	log.SetFormatter(&log.TextFormatter{
		EnvironmentOverrideColors: true,
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           time.RFC3339,
	})
	log.SetLevel(log.DebugLevel)
	routes.UseRoute(e)

	database.Database.Connect()
	go initSitemapGenerator()
	sitemap.SetCronForSitemap()
	latest_bar.SetLatestCron()
	service.SetCronForIndexMap()

	latest_bar.InitLatestComments()
	latest_bar.InitLatestSearches()
}

func initSitemapGenerator() {
	config.ApiConfig.Sitemaps = sitemap.GenerateNumbersSitemap()
	config.ApiConfig.Sitemap = sitemap.GenerateBaseSitemap()
	config.ApiConfig.SitemapMain = sitemap.GenerateMainSitemap()
}
