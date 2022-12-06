package companyRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) CreateUserCompany(UserId int, company *models.Company) (*models.Company, error) {

	user := models.User{}

	if err := r.db.Create(company).
		Table("users").Model(&user).Where("id=?", UserId).
		Updates(models.User{CompanyId: &company.Id}).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (r *CompanyRepository) GetCompanyByUserId(id int) (*models.Company, error) {

	user := &models.User{}

	if err := r.db.Preload("Company").Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user.Company, nil
}

func (r *CompanyRepository) CheckingCompanyOfUserIsExits(userId int) bool {

	var count int64

	if err := r.db.Where("id=?", userId).Where("company_id IS NOT NULL").Find(&models.User{}).Count(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (r *CompanyRepository) GetCompanyActivity() ([]models.CompanyActivity, error) {

	companyActivities := []models.CompanyActivity{}

	err := r.db.Find(&companyActivities).Error

	return companyActivities, err
}
func (r *CompanyRepository) GetCompanyByID(id int) (*models.Company, error) {

	company := &models.Company{}

	err := r.db.Preload("Location").Preload("CompanyActivity").Where("id=?", id).First(company).Error

	return company, err
}
func (r *CompanyRepository) UpdateCompany(id int, company *models.Company) error {

	err := r.db.Where("id=?", id).UpdateColumns(company).Error

	return err
}
func (r *CompanyRepository) GetCompany() ([]models.Company, error) {
	var company []models.Company
	if err := r.db.Find(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}
func (r *CompanyRepository) DeleteCompany(id int) error {
	if err := r.db.Delete(&models.Company{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
