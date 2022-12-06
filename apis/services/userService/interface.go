package userService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type UserServiceInterface interface {
	GetProfileById(id int) (*models.User, *response.CustomError)
	UpdateUserProfile(id int, user *models.User) *response.CustomError
	ChangePassword(id int, changePass *dto.ChangePassword) *response.CustomError
}
