package uploadRepository

import "mime/multipart"

type UpLoadRepositoryInterface interface {
	UploadImage(file *multipart.FileHeader, folder string) (string, error)
}
