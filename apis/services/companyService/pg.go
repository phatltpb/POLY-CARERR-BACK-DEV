package companyService

import (
	"errors"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/companyRepository"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type CompanyService struct {
	companyRepository companyRepository.CompanyRepositoryInterface
}

func Init(repositories *repositories.Repositories) *CompanyService {
	return &CompanyService{
		companyRepository: repositories.CompanyRepository,
	}
}

func (r *CompanyService) CreateUserCompany(company *models.Company, idUser int) (int, *response.CustomError) {

	if r.companyRepository.CheckingCompanyOfUserIsExits(idUser) {
		return 0, &response.CustomError{
			Error: errors.New(validationMessage.CompanyIsExits),
		}
	}

	validation := helper.InitValidation().Require("name", company.Name).
		Require("tax_code", company.TaxCode)

	if !validation.IsValid {
		return 0, &response.CustomError{
			Code:       mystatus.UnprocessableEntity,
			ErrorField: validation.ErrorField,
		}
	}

	rs, err := r.companyRepository.CreateUserCompany(idUser, company)

	if err != nil {
		return 0, &response.CustomError{
			Error: err,
		}
	}

	return rs.Id, nil
}

func (r *CompanyService) GetCompanyActivity() ([]models.CompanyActivity, *response.CustomError) {

	companyActivities, err := r.companyRepository.GetCompanyActivity()

	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}
	}
	return companyActivities, nil

}
func (r *CompanyService) UpdateCompany(company *models.Company, idUser int) *response.CustomError {

	var (
		myCompany *models.Company
		err       error
	)

	if myCompany, err = r.companyRepository.GetCompanyByUserId(idUser); err != nil || myCompany == nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}

	if err := r.companyRepository.UpdateCompany(myCompany.Id, company); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}

	return nil

}

func (r *CompanyService) GetCompanies() ([]models.Company, *response.CustomError) {
	company, err := r.companyRepository.GetCompany()
	if err != nil {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	return company, nil
}

func (r *CompanyService) GetCompaniesById(id int) (*models.Company, *response.CustomError) {
	company, err := r.companyRepository.GetCompanyByID(id)
	if err != nil {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}
	return company, nil
}

func (r *CompanyService) DeleteCompany(id int) *response.CustomError {
	err := r.companyRepository.DeleteCompany(id)
	if err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	return nil
}
