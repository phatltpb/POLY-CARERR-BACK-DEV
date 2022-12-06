package studentService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type StudentServiceInterface interface {
	GetProfile(id int) (*models.Students, *response.CustomError)
	UpdateStudentProfile(id int, student *models.Students) *response.CustomError

	CreateOrUpdateStudentDetail(idStudent int, profile *models.StudentProfile) *response.CustomError

	CreateStudentEducation(education *models.StudentEducation) *response.CustomError
	UpdateStudentEducation(id int, education *models.StudentEducation) *response.CustomError
	DeleteStudentEducation(id int) *response.CustomError
	ChangePassword(id int, changePass *dto.ChangePassword)
}
