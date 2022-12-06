package v1

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type CompanyCtl struct {
	services *services.Service
}

func InitCompanyCtl(service *services.Service) *CompanyCtl {
	return &CompanyCtl{
		services: service,
	}
}

func (ctl *CompanyCtl) GetCompanyActivity(c echo.Context) error {

	companyActivities, err := ctl.services.CompanyService.GetCompanyActivity()

	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, companyActivities)

}
func (ctl *CompanyCtl) CreateCompany(c echo.Context) error {

	company := &models.Company{}

	if err := c.Bind(company); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	id, err := ctl.services.CompanyService.CreateUserCompany(company, jwtToken.Id)

	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, id)
}

func (ctl *CompanyCtl) UpdateCompany(c echo.Context) error {

	company := &models.Company{}

	if err := c.Bind(company); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	if err := ctl.services.CompanyService.UpdateCompany(company, jwtToken.Id); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, nil)
}

func (ctl *CompanyCtl) GetCompany(c echo.Context) error {
	company, err := ctl.services.CompanyService.GetCompanies()

	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, company)
}

func (ctl *CompanyCtl) GetCompanyById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(c, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.ApiNotFound),
		})
	}
	company, err1 := ctl.services.CompanyService.GetCompaniesById(id)
	if err1 != nil {
		return response.Error(c, err1)
	}
	return response.Success(c, company)
}
