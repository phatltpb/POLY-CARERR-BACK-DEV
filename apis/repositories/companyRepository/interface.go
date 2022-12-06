package companyRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type CompanyRepositoryInterface interface {
	CreateUserCompany(UserId int, company *models.Company) (*models.Company, error)
	GetCompanyByUserId(id int) (*models.Company, error)
	CheckingCompanyOfUserIsExits(id int) bool
	GetCompanyActivity() ([]models.CompanyActivity, error)
	GetCompanyByID(id int) (*models.Company, error)
	UpdateCompany(id int, company *models.Company) error
	GetCompany() ([]models.Company, error)
	DeleteCompany(id int) error
}
