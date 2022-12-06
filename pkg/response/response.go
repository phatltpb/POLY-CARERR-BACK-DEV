package response

import (
	"github.com/labstack/echo/v4"
	myStatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
)

type responseSuccess struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type responseSuccessWithMeta struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
}

type responseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ErrorField struct {
	Field string   `json:"field"`
	Code  []string `json:"message"`
}
type responseErrorWithField struct {
	Code       string       `json:"code"`
	Message    string       `json:"message"`
	Error      string       `json:"error"`
	ErrorField []ErrorField `json:"error_field,omitempty"`
}

func Success(c echo.Context, data interface{}) error {

	codeMapper := myStatus.CodeMapper[myStatus.Success]

	return c.JSON(codeMapper.StatusCode, responseSuccess{
		Data:    data,
		Code:    codeMapper.Code,
		Message: codeMapper.Message,
	})
}

func SuccessWithMeta(c echo.Context, data interface{}, meta interface{}) error {

	codeMapper := myStatus.CodeMapper[myStatus.Success]

	return c.JSON(codeMapper.StatusCode, responseSuccessWithMeta{
		Data:    data,
		Code:    codeMapper.Code,
		Message: codeMapper.Message,
		Meta:    meta,
	})
}

func Error(c echo.Context, err *CustomError) error {
	codeMapper, exists := myStatus.CodeMapper[err.Code]
	if !exists {
		codeMapper = myStatus.CodeMapper[myStatus.BadRequest]
	}

	res := responseErrorWithField{
		Code:    codeMapper.Code,
		Message: codeMapper.Message,
	}

	if err.ErrorField != nil {
		res.ErrorField = err.ErrorField
	}

	if err.Error != nil && err.Error.Error() != "" {
		res.Error = err.Error.Error()
	} else {
		res.Error = "Api error!!!"
	}

	return c.JSON(codeMapper.StatusCode, res)
}
