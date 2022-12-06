package uploadService

import (
	"mime/multipart"

	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type UploadServiceInterface interface {
	UploadImage(file *multipart.FileHeader) (string, *response.CustomError)
	UploadFile(file *multipart.FileHeader) (string, *response.CustomError)
}
