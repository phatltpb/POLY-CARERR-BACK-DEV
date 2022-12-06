package studentCvRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type StudentCvRepositoryInterface interface {
	CreateStudentCv(studentCv *models.StudentCV) (*models.StudentCV, error)
	GetCvByIdStudent(id int) ([]models.StudentCV, error)
}
