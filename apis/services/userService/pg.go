package userService

import (
	"errors"

	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/userRepository"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type UserService struct {
	UserRepo userRepository.UserRepositoryInterface
}

func Init(repositories *repositories.Repositories) *UserService {
	return &UserService{
		UserRepo: repositories.UserRepository,
	}
}

func (uS *UserService) GetProfileById(id int) (*models.User, *response.CustomError) {
	if id == 0 {
		return nil, &response.CustomError{
			Error: errors.New(validationMessage.UserIsNotExits),
			Code:  mystatus.BadRequest,
		}
	}

	users, err := uS.UserRepo.GetUserById(id)

	if err != nil {
		return nil, &response.CustomError{
			Error: err,
			Code:  mystatus.BadRequest,
		}
	}

	return users, nil
}

func (Us *UserService) UpdateUserProfile(id int, user *models.User) *response.CustomError {
	if err := Us.UserRepo.UpdateUser(id, user); err != nil {
		return &response.CustomError{
			Error: err,
			Code:  mystatus.BadRequest,
		}
	}
	return nil
}

func (Us *UserService) ChangePassword(id int, changePass *dto.ChangePassword) *response.CustomError {

	user, err := Us.UserRepo.GetUserById(id)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	if err := helper.CheckPasswordHash(user.Password, changePass.CurrentPassword); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.PasswordWrong),
		}
	}

	var newUser models.User
	newUser.Password, _ = helper.Hash(changePass.NewPassword)
	if err := Us.UserRepo.UpdateUser(id, &newUser); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return nil
}
