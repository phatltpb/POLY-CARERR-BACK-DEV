package response

type CustomError struct {
	Error      error
	Code       string
	ErrorField []ErrorField
}
