package authservice

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type AuthInterface interface {
	Login(res *dto.LoginDto) (string, *response.CustomError)
	Register(res *models.Students) (string, *response.CustomError)
	LoginAdmin(res *dto.LoginDto) (string, *response.CustomError)
	RegisterAdmin(res *models.User) (string, *response.CustomError)
	VerifyAdminRegister(token string) (string, *response.CustomError)
	VerifyRegister(token string) (string, *response.CustomError)
	// student
	ForgotStudentPassword(email string) (string, *response.CustomError)
	ResetStudentPassword(token string, password string) *response.CustomError

	// Admin
	ForGotAdminPassword(email string) (string, *response.CustomError)
	ResetAdminPassword(token string, password string) *response.CustomError
}
