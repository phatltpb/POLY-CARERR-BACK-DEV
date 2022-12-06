package uploadService

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/uploadRepository"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type UploadService struct {
	UploadRepo uploadRepository.UpLoadRepositoryInterface
}

func Init(repositories *repositories.Repositories) *UploadService {
	return &UploadService{
		UploadRepo: repositories.UploadRepository,
	}
}

func (s *UploadService) UploadImage(file *multipart.FileHeader) (string, *response.CustomError) {
	upload, err := s.UploadRepo.UploadImage(file, "polycareer/image")

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.UploadError),
		}
	}
	return upload, nil
}

func (s *UploadService) UploadFile(file *multipart.FileHeader) (string, *response.CustomError) {
	upload, err := s.UploadRepo.UploadImage(file, "polycareer/cv")

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.UploadError),
		}
	}
	if !strings.Contains(file.Filename, "zip") || !strings.Contains(file.Filename, "pdf") {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.UploadError),
		}
	}
	return upload, nil
}
