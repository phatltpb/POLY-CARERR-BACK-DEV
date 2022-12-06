package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type UploadCtl struct {
	services *services.Service
}

func InitUpload(services *services.Service) *UploadCtl {
	return &UploadCtl{
		services: services,
	}
}
func (ctl *UploadCtl) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	createFile, err1 := ctl.services.UploadService.UploadImage(file)
	if err1 != nil {
		return response.Error(c, err1)
	}

	return response.Success(c, createFile)
}
func (ctl *UploadCtl) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	createFile, er := ctl.services.UploadService.UploadFile(file)
	if er != nil {
		return response.Error(c, er)
	}
	return response.Success(c, createFile)
}
