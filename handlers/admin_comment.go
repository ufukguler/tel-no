package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"telno/models"
	"telno/service"
)

func FindCommentById(c echo.Context) error {
	var dto models.FindByCommentIdRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	phoneNumber, comment, err := service.FindCommentById(dto.ID)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	return e.ResponseOk(models.SingleComment{
		Id:          comment.Id,
		Comment:     comment.Comment,
		PhoneNumber: phoneNumber,
		CommentType: comment.CommentType,
		Updated:     comment.Updated,
		CreatedDate: comment.CreatedDate,
		UpdatedDate: comment.UpdatedDate,
	}, c)
}

func DeleteCommentById(c echo.Context) error {
	var dto models.FindByCommentIdRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := service.DeleteCommentById(dto.ID); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	return e.ResponseOk(true, c)
}

func UpdateCommentById(c echo.Context) error {
	var dto models.UpdateByCommentIdRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := service.UpdateCommentById(dto.ID, dto.Comment, dto.CommentType); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	return e.ResponseOk(true, c)
}

func QuickUpdateCommentById(c echo.Context) error {
	var dto models.QuickUpdateByCommentIdRequestDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := service.QuickUpdateCommentById(dto.ID); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	return e.ResponseOk(true, c)
}
