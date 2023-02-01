package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"telno/config"
	"telno/models"
)

const recaptchaUrl = "https://www.google.com/recaptcha/api/siteverify"

func CheckRecaptcha(c echo.Context, responseKey string) bool {
	client := &http.Client{}
	data := url.Values{}
	data.Set("secret", config.ApiConfig.CaptchaSecret)
	data.Set("response", responseKey)
	data.Set("remoteip", c.RealIP())

	r, _ := http.NewRequest(http.MethodPost, recaptchaUrl, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		log.Errorf("http call error: %s", err.Error())
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("body read error: %s", err.Error())
		return false
	}

	var recaptcha models.Recaptcha
	err = json.Unmarshal(body, &recaptcha)
	if err != nil {
		log.Errorf("json unmarshal error: %s", err.Error())
		return false
	}
	if recaptcha.Success {
		return true
	}
	log.Errorf("Recaptcha | ErrorCodes: %s", recaptcha.ErrorCodes)
	return false
}
