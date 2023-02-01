package handlers

import (
	"bytes"
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"strings"
	"telno/models"
	"telno/service/image_service"
)

func GetImagePhoneNumber(c echo.Context) error {
	var dto models.PhoneNumberImageRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	// CHECK IF PHONE NUMBER STARTS WITH ZERO
	if string([]rune(dto.PhoneNumber)[0]) == "0" {
		dto.PhoneNumber = dto.PhoneNumber[1:]
	}
	// CHECK IF PHONE NUMBER IS VALID
	trimmed := strings.TrimSpace(dto.PhoneNumber)
	if len(trimmed) == 0 {
		return e.ResponseError(errors.New("phone number can not be empty"), c)
	}
	firstNumber := string([]rune(trimmed)[0])
	if len(trimmed) == 10 && (firstNumber == "2" || firstNumber == "3" || firstNumber == "5" || firstNumber == "8") {
		//ok
	} else if len(trimmed) == 7 && trimmed[:3] == "444" {
		//ok
	} else {
		return e.ResponseError(errors.New("invalid phone number"), c)
	}

	image, err := generateImage(trimmed)
	if err != nil {
		return err
	}
	return c.Blob(200, "image/png", image)
}

func generateImage(trimmed string) ([]byte, error) {

	// GENERATE IMAGE
	img, err := image_service.TextOnImg(trimmed)
	if err != nil {
		log.Error("[ERROR] TextOnImg: ", err.Error())
		return nil, err
	}
	// PREPARE IMAGE
	buf := new(bytes.Buffer)
	if err = jpeg.Encode(buf, img, nil); err != nil {
		log.Error("[ERROR] jpeg encode: ", err.Error())
		return nil, err
	}

	return buf.Bytes(), nil
}
