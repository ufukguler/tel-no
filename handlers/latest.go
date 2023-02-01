package handlers

import (
	"github.com/labstack/echo/v4"
	"telno/models"
	"telno/service/latest_bar"
)

func GetLatestUpdates(c echo.Context) error {
	latest := models.Latest{
		LatestComments: latest_bar.GetLatestComments(),
		LatestSearches: latest_bar.GetLatestSearches(),
	}
	return e.ResponseOk(latest, c)
}
