package mystatus

import "net/http"

type CodeMapperStruct struct {
	StatusCode int
	Code       string
	Message    string
}

var (
	// "200000"
	Success string = "200000"
	//400000
	BadRequest string = "400000"
	//500000
	InternalServerError string = "500000"
	// 401000
	Unauthorized string = "401000"
	// 401001
	TokenExpired string = "401001"
	// 401003
	TokenMissing string = "401003"
	// 401002
	TokenInvalid string = "401002"
	// 422000
	UnprocessableEntity string = "422000"
	PageNotFound        string = "404000"
	//428000
	UnUploadFile string = "403000"
	//403   not authentication
)

var CodeMapper map[string]CodeMapperStruct = map[string]CodeMapperStruct{
	Success: {
		StatusCode: http.StatusOK,
		Code:       "200",
		Message:    "Success!",
	},
	BadRequest: {
		StatusCode: http.StatusBadRequest,
		Code:       "400000",
		Message:    "Bad request!",
	},
	InternalServerError: {
		StatusCode: http.StatusInternalServerError,
		Code:       "500000",
		Message:    "Server error!",
	},
	Unauthorized: {
		StatusCode: http.StatusUnauthorized,
		Code:       "401000",
		Message:    "Unauthorized!",
	},
	TokenExpired: {
		StatusCode: http.StatusUnauthorized,
		Code:       "401001",
		Message:    "Token expired!",
	},
	TokenInvalid: {
		StatusCode: http.StatusUnauthorized,
		Code:       "401001",
		Message:    "Token invalid!",
	},
	TokenMissing: {
		StatusCode: http.StatusUnauthorized,
		Code:       "401001",
		Message:    "Token missing!",
	},
	UnprocessableEntity: {
		StatusCode: http.StatusUnprocessableEntity,
		Code:       "422000",
		Message:    "Unprocessable entity!",
	},
	PageNotFound: {
		StatusCode: http.StatusNotFound,
		Code:       "404000",
		Message:    "Page not found!",
	},
}
