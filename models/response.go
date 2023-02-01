package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Request interface{} `json:"requestUrl,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (resp *Response) ResponseCreated(data interface{}, c echo.Context) error {
	return c.JSON(http.StatusCreated,
		Response{
			Code:    http.StatusCreated,
			Message: "Created",
			Data:    data,
		})
}

func (resp *Response) ResponseOk(data interface{}, c echo.Context) error {
	return c.JSON(http.StatusOK,
		Response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		})
}

func (resp *Response) ResponseOkEmpty(c echo.Context) error {
	return c.JSON(http.StatusOK,
		Response{
			Code:    http.StatusOK,
			Message: "Ok",
		})
}

func (resp *Response) ResponseError(err error, c echo.Context) error {
	return c.JSON(http.StatusBadRequest,
		Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
}

func (resp *Response) ResponseNotFound(err error, c echo.Context) error {
	return c.JSON(http.StatusNotFound,
		Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Request: c.Request().URL.RequestURI(),
		})
}

func (resp *Response) ResponseUnauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized,
		Response{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		})
}

func (resp *Response) ResponseForbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden,
		Response{
			Code:    http.StatusForbidden,
			Message: "forbidden",
		})
}

func (resp *Response) ResponseFatal(err error, c echo.Context) error {
	return c.JSON(http.StatusInternalServerError,
		Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
}
