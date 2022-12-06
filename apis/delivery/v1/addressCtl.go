package v1

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type AddressCtl struct {
	services *services.Service
}

func InitAddressCtl(service *services.Service) *AddressCtl {
	return &AddressCtl{
		services: service,
	}
}

func (ctl *AddressCtl) GetAddress(c echo.Context) error {

	data, err := ctl.services.AddressService.GetAddress()

	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
			Code:  mystatus.BadRequest,
		})
	}
	return response.Success(c, data)
}
