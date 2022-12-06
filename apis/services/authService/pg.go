package authservice

import (
	"errors"
	"time"

	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	repository "github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/userRepository"
	"github.com/tuongnguyen1209/poly-career-back/config"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	"github.com/tuongnguyen1209/poly-career-back/pkg/role"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type AuthService struct {
	StudentRepo repository.StudentRepositoryInterface
	UserRepo    userRepository.UserRepositoryInterface
}

func Init(repositories *repositories.Repositories) *AuthService {

	return &AuthService{
		StudentRepo: repositories.StudentRepository,
		UserRepo:    repositories.UserRepository,
	}
}

func (s *AuthService) Login(res *dto.LoginDto) (string, *response.CustomError) {

	validation := helper.InitValidation().Require("email", res.Email).Require("password", res.Password)

	if !validation.IsValid {
		return "", &response.CustomError{
			Code:       mystatus.UnprocessableEntity,
			ErrorField: validation.ErrorField,
		}
	}
	student, err := s.StudentRepo.GetUserByEmail(res.Email)

	if err != nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.LoginFalse),
			Code:  mystatus.UnprocessableEntity,
		}
	}
	if !student.IsActive {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.UnActive),
			Code:  mystatus.UnprocessableEntity,
		}
	}
	if err := helper.CheckPasswordHash(student.Password, res.Password); err != nil {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: errors.New(validationMessage.LoginFalse),
		}
	}

	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.GetConfig().Jwt.Expires)

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(student.Id),
			Role:   student.Role,
			Action: pkg.JwtAction.Login,
		},
		duration)

	if err != nil {

		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: errors.New(validationMessage.LoginFalse),
		}
	}

	return token, nil
}

func (s *AuthService) Register(res *models.Students) (string, *response.CustomError) {

	validation := helper.InitValidation().Require("email", res.Email).Require("password", res.Password).IsEmail("email", res.Email)

	if s.StudentRepo.CheckingEmail(res.Email) {
		validation.SetError("email", validationMessage.IsExits)
	}

	if !validation.IsValid {
		return "", &response.CustomError{
			Code:       mystatus.UnprocessableEntity,
			ErrorField: validation.ErrorField,
		}
	}

	newPassword, err := helper.Hash(res.Password)
	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: err,
		}
	}

	res.Password = newPassword
	res.Role = role.Student
	student, err := s.StudentRepo.CreateStudent(res)

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: err,
		}
	}

	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.GetConfig().Jwt.Expires)

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(student.Id),
			Role:   student.Role,
			Action: pkg.JwtAction.Verify,
		}, duration)

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.LoginFalse),
		}
	}

	return token, nil
}
func (s *AuthService) VerifyRegister(token string) (string, *response.CustomError) {

	jwt := pkg.JwtConfig{}

	if token == "" {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenMissing),
			Code:  mystatus.TokenMissing,
		}
	}
	claims, err := jwt.Decode(token)
	if err != nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims == nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	res := s.StudentRepo.UpdateActive(claims.Id)
	if !res {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.BadRequest,
		}
	}

	return "", nil
}

func (s *AuthService) LoginAdmin(res *dto.LoginDto) (string, *response.CustomError) {

	if res.Email == "" || res.Password == "" {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: errors.New(validationMessage.LoginFalse),
		}
	}

	user, err := s.UserRepo.GetUserByEmail(res.Email)

	if err != nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.LoginFalse),
			Code:  mystatus.UnprocessableEntity,
		}
	}

	if !user.IsActive {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.UnActive),
			Code:  mystatus.UnprocessableEntity,
		}
	}

	if err := helper.CheckPasswordHash(user.Password, res.Password); err != nil {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: errors.New(validationMessage.LoginFalse),
		}
	}

	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.GetConfig().Jwt.Expires)

	var action string = pkg.JwtAction.LoginEmployer

	if user.Role == role.Admin {
		action = pkg.JwtAction.LoginEmployer
	}

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(user.Id),
			Role:   user.Role,
			Action: action,
		},
		duration)

	if err != nil {

		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: err,
		}
	}

	return token, nil
}

func (s *AuthService) RegisterAdmin(res *models.User) (string, *response.CustomError) {

	validation := helper.InitValidation().
		Require("email", res.Email).
		Require("password", res.Password).IsEmail("email", res.Email)

	if s.UserRepo.CheckingEmail(res.Email) {
		validation.SetError("email", validationMessage.IsExits)
	}

	if !validation.IsValid {
		return "", &response.CustomError{
			Code:       mystatus.UnprocessableEntity,
			ErrorField: validation.ErrorField,
		}
	}

	newPassword, err := helper.Hash(res.Password)
	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: err,
		}
	}

	res.Password = newPassword
	res.Role = role.Employer

	user, err := s.UserRepo.CreateUser(res)

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: err,
		}
	}
	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.GetConfig().Jwt.Expires)

	var action string = pkg.JwtAction.LoginEmployer

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(user.Id),
			Role:   user.Role,
			Action: action,
		},
		duration)
	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.UnprocessableEntity,
			Error: err,
		}
	}
	return token, nil

}
func (s *AuthService) VerifyAdminRegister(token string) (string, *response.CustomError) {

	jwt := pkg.JwtConfig{}

	if token == "" {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenMissing),
			Code:  mystatus.TokenMissing,
		}
	}
	claims, err := jwt.Decode(token)
	if err != nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims == nil {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	res := s.UserRepo.UpdateActive(claims.Id)
	if !res {
		return "", &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.BadRequest,
		}
	}

	return "", nil
}

// / forgot student password request
func (s *AuthService) ForgotStudentPassword(email string) (string, *response.CustomError) {
	config := config.GetConfig()
	student, err := s.StudentRepo.GetUserByEmail(email)
	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.Jwt.Expires)

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(student.Id),
			Role:   student.Role,
			Action: pkg.JwtAction.ResetPassword,
		}, duration)

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.ApiNotFound),
		}
	}

	return token, nil
}

func (s *AuthService) ResetStudentPassword(token string, password string) *response.CustomError {
	jwt := pkg.JwtConfig{}

	if token == "" {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenMissing),
			Code:  mystatus.TokenMissing,
		}
	}
	claims, err := jwt.Decode(token)
	if err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims == nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims.Action != pkg.JwtAction.ResetPassword {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims.Role != role.Student {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	var student models.Students
	student.Password, _ = helper.Hash(password)
	if err := s.StudentRepo.UpdateStudent(&student, claims.Id); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: err,
		}
	}
	return nil
}

// forget employer password
func (s *AuthService) ForGotAdminPassword(email string) (string, *response.CustomError) {
	config := config.GetConfig()
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	jwt := pkg.JwtConfig{}
	duration := time.Hour * 24 * time.Duration(config.Jwt.Expires)

	token, err := jwt.Encode(
		pkg.JwtToken{
			Id:     int(user.Id),
			Role:   user.Role,
			Action: pkg.JwtAction.ResetPassword,
		}, duration)

	if err != nil {
		return "", &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.ApiNotFound),
		}
	}

	return token, nil
}

func (s *AuthService) ResetAdminPassword(token string, password string) *response.CustomError {
	jwt := pkg.JwtConfig{}
	if token == "" {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenMissing),
			Code:  mystatus.TokenMissing,
		}
	}
	claims, err := jwt.Decode(token)
	if err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims == nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims.Action != pkg.JwtAction.ResetPassword {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	if claims.Role != role.Employer {
		return &response.CustomError{
			Error: errors.New(validationMessage.TokenInvalid),
			Code:  mystatus.TokenInvalid,
		}
	}
	var user models.User
	user.Password, _ = helper.Hash(password)
	if err := s.UserRepo.UpdateUser(claims.Id, &user); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: err,
		}
	}
	return nil
}

//
