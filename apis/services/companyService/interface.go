package companyService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type CompanyServiceInterface interface {
	CreateUserCompany(company *models.Company, idUser int) (int, *response.CustomError)
	GetCompanyActivity() ([]models.CompanyActivity, *response.CustomError)
	UpdateCompany(company *models.Company, idUser int) *response.CustomError
	GetCompaniesById(id int) (*models.Company, *response.CustomError)
	GetCompanies() ([]models.Company, *response.CustomError)
	DeleteCompany(id int) *response.CustomError
}
