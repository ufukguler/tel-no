package handlers

import (
	"github.com/gobeam/mongo-go-pagination"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"telno/model_entity"
	"telno/models"
	"telno/service"
)

func FindPhoneNumbersByPageable(c echo.Context) error {
	var dto models.PageableRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	limit, _ := strconv.ParseInt(dto.Limit, 10, 64)
	page, _ := strconv.ParseInt(dto.Page, 10, 64)
	pageData, err := service.FindPhoneNumberPageable(limit, page, dto.PhoneNumber)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	pageableResponse := initPhoneNumberPageableResponse(pageData)

	return e.ResponseOk(models.PageableResponseDTO{
		Data:     pageableResponse,
		Pageable: pageData.Pagination,
	}, c)
}

func FindPhoneNumbersCommentUncheckedByPageable(c echo.Context) error {
	var dto models.PageableRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	limit, _ := strconv.ParseInt(dto.Limit, 10, 64)
	page, _ := strconv.ParseInt(dto.Page, 10, 64)
	pageData, err := service.FindPhoneNumberByCommentUncheckedPageable(limit, page, dto.PhoneNumber)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	pageableResponse := initPhoneNumberPageableResponse(pageData)

	return e.ResponseOk(models.PageableResponseDTO{
		Data:     pageableResponse,
		Pageable: pageData.Pagination,
	}, c)
}

func initPhoneNumberPageableResponse(pageData *mongopagination.PaginatedData) []models.CrawledUrlAdmin {
	numberList := make([]models.CrawledUrlAdmin, 0)
	for _, raw := range pageData.Data {
		var phoneNumber models.CrawledUrlAdmin
		if marshallErr := bson.Unmarshal(raw, &phoneNumber); marshallErr == nil {
			numberList = append(numberList, phoneNumber)
		}
	}
	return numberList
}

func FindByPhoneNumberOnlyFalseComments(c echo.Context) error {
	number := c.Param("number")
	crawledUrl, err := service.FindByPhoneNumber(number)
	if err != nil {
		return e.ResponseError(err, c)
	}

	comments := make([]model_entity.Comment, 0)
	for i := range crawledUrl.Comments {
		if !crawledUrl.Comments[i].Updated {
			comments = append(comments, crawledUrl.Comments[i])
		}
	}
	crawledUrl.Comments = comments
	return e.ResponseOk(crawledUrl, c)
}
