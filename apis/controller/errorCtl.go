package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

func HandleNotFound(c echo.Context) error {

	return response.Error(c, &response.CustomError{
		Error: errors.New(validationMessage.ApiNotFound),
		Code:  mystatus.PageNotFound,
	})
}

func CustomHTTPErrorHandler(err error, c echo.Context) error {
	c.Logger().Error(err)
	return response.Error(c, &response.CustomError{
		Code: mystatus.InternalServerError,
	})
}
