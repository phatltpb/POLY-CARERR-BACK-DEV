package userRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type UserRepositoryInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	CheckingEmail(email string) bool
	CheckingUserAdmin(id int) bool
	UpdateActive(id int) bool
	UpdateUser(id int, user *models.User) error
}
