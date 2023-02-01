package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"strings"
	"telno/model_entity"
	"telno/models"
	"telno/service"
	"telno/service/latest_bar"
)

var e models.Response

func FindByPhoneNumber(c echo.Context) error {
	var dto models.FindByPhoneNumberRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	dto.PhoneNumber = strings.TrimSpace(dto.PhoneNumber)
	if len(strings.TrimSpace(dto.PhoneNumber)) == 0 {
		return e.ResponseError(errors.New("phone number can not be empty"), c)
	}

	// CHECK IF PHONE NUMBER STARTS WITH ZERO
	if string([]rune(dto.PhoneNumber)[0]) == "0" {
		dto.PhoneNumber = dto.PhoneNumber[1:]
	}
	// CHECK IF PHONE NUMBER IS VALID
	if isPhoneNumberValid(dto.PhoneNumber) == false {
		return e.ResponseError(errors.New("invalid phone number"), c)
	}

	number, err := service.FindByPhoneNumber(dto.PhoneNumber)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			latest_bar.UpdateLatestSearches(dto.PhoneNumber)
			return e.ResponseOk(initNotFoundNumber(dto.PhoneNumber), c)
		}
		log.Errorf("MongoDB Error: %s", err.Error())
		return e.ResponseNotFound(errors.New("error"), c)
	}

	// INIT RESPONSE
	respArr := initSimilarNumbers(number.PhoneNumber)
	reading := ""
	if len(number.PhoneNumber) == 10 {
		reading = service.GetWritingOfPhoneNumber10Digit(number.PhoneNumber)
	}
	if len(number.PhoneNumber) == 7 && strings.HasPrefix(number.PhoneNumber, "4") {
		reading = service.GetWritingOfPhoneNumber7Digit(number.PhoneNumber)
	}
	commentsUpdated := make([]model_entity.Comment, 0)
	for i := range number.Comments {
		if number.Comments[i].Updated {
			commentsUpdated = append(commentsUpdated, number.Comments[i])
		}
	}
	crawledUrlDTO := models.CrawledUrlResponseDTO{
		PhoneNumber:     number.PhoneNumber,
		Comments:        commentsUpdated,
		Reading:         reading,
		Operator:        initOperator(number.PhoneNumber),
		PhoneType:       initPhoneType(number.PhoneNumber),
		PhoneRegion:     initPhoneRegion(number.PhoneNumber),
		ReviewCount:     number.ViewCount,
		SimilarNumbers:  respArr,
		AggregateRating: service.GenerateAggregateRating(number),
	}

	go func() {
		if strings.Contains(c.Request().Header.Get("origin"), "localhost:") {
			ip := c.Request().Header.Get("X-Real-IP")
			service.UpdateRequestCount(number.Id, ip)
		}
	}()
	latest_bar.UpdateLatestSearches(dto.PhoneNumber)
	return e.ResponseOk(crawledUrlDTO, c)
}

func AddComment(c echo.Context) error {
	var dto models.AddCommentRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	trimmed := strings.TrimSpace(dto.PhoneNumber)
	if len(trimmed) == 0 {
		return e.ResponseError(errors.New("phone number can not be empty"), c)
	}

	// CHECK IF PHONE NUMBER STARTS WITH ZERO
	if string([]rune(dto.PhoneNumber)[0]) == "0" {
		dto.PhoneNumber = dto.PhoneNumber[1:]
	}
	// CHECK IF PHONE NUMBER IS VALID
	if isPhoneNumberValid(trimmed) == false {
		return e.ResponseError(errors.New("invalid phone number"), c)
	}

	// CHECK IF COMMENT TYPE IS VALID
	if err := checkCommentType(dto.CommentType); err != nil {
		return e.ResponseError(err, c)
	}

	// CHECK CAPTCHA
	if service.CheckRecaptcha(c, dto.ResponseKey) == false {
		return e.ResponseError(errors.New("invalid captcha"), c)
	}

	crawledUrl, err := service.FindByPhoneNumber(dto.PhoneNumber)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			comments := initCommentsForFirstTimeCrawledUrl(dto)
			savedCrawledUrl := initFirstTimeCrawledUrl(dto, comments)
			if err = service.CreatePhoneNumber(savedCrawledUrl); err != nil {
				log.Errorf("CreatePhoneNumber error : %s", err.Error())
				return e.ResponseError(errors.New("an error occurred"), c)
			}
			latest_bar.UpdateLatestComments(dto.PhoneNumber, dto.Comment)
			return e.ResponseOk(true, c)
		}
		log.Errorf("MongoDB Error: %s", err.Error())
		return e.ResponseError(errors.New("an error occurred"), c)
	}
	crawledUrl = addNewCommentToArray(dto, crawledUrl)
	if err = service.UpdatePhoneNumber(crawledUrl); err != nil {
		log.Errorf("UpdatePhoneNumber | MongoDB Error: %s", err.Error())
		return e.ResponseError(errors.New("an error occurred"), c)

	}

	latest_bar.UpdateLatestComments(dto.PhoneNumber, dto.Comment)
	return e.ResponseOk(true, c)
}

func isPhoneNumberValid(trimmed string) bool {

	// CHECK IF PHONE NUMBER IS VALID
	firstNumber := string([]rune(trimmed)[0])
	if len(trimmed) == 10 && (firstNumber == "2" || firstNumber == "3" || firstNumber == "5" || firstNumber == "8") {
		//ok
	} else if len(trimmed) == 7 && trimmed[:3] == "444" {
		//ok
	} else {
		return false
	}
	return true
}
