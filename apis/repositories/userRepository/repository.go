package userRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/role"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user *models.User) (*models.User, error) {
	rs := u.db.Create(user)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return user, nil
}

func (u UserRepository) GetUserByEmail(email string) (*models.User, error) {

	user := models.User{}

	rs := u.db.Where("email=?", email).First(&user)

	if rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (u UserRepository) GetUserById(id int) (*models.User, error) {

	user := &models.User{}

	rs := u.db.Preload("Company.Location").Preload("Company.CompanyActivity").Where("id=?", id).First(user)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return user, nil
}

func (u UserRepository) CheckingEmail(email string) bool {

	var count int64
	if u.db.Model(&models.User{}).Where("email=?", email).Count(&count).Error != nil {
		return false
	}
	return count > int64(0)
}

func (u UserRepository) CheckingUserAdmin(id int) bool {
	user := models.User{}

	if u.db.Where("id=?", id).First(&user).Error != nil {
		return false
	}

	return user.Role == role.Admin
}
func (u UserRepository) UpdateActive(id int) bool {
	if err := u.db.Where("id = ?", id).Updates(&models.User{IsActive: true}).Error; err != nil {
		return false
	}
	return true
}

func (u UserRepository) UpdateUser(id int, user *models.User) error {
	if err := u.db.Where("id = ? ", id).UpdateColumns(&user).Error; err != nil {
		return err
	}

	return nil
}
