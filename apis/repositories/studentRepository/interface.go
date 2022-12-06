package repository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type StudentRepositoryInterface interface {
	GetUserByEmail(email string) (*models.Students, error)
	CreateStudent(student *models.Students) (*models.Students, error)
	GetStudentById(id int) (*models.Students, error)
	CheckingEmail(email string) bool
	UpdateActive(id int) bool
	UpdateStudent(student *models.Students, id int) error

	GetStudentDetailByStudentId(studentId int) (*models.StudentProfile, error)
	CreateStudentDetail(studentProfile *models.StudentProfile) error
	UpdateStudentDetail(id int, studentProfile *models.StudentProfile) error
	GetStudentEducationById(id int) (*models.StudentEducation, error)
	CreateStudentEducation(education *models.StudentEducation) error
	UpdateStudentEducation(id int, education *models.StudentEducation) error
	DeleteStudentEducation(id int) error
}
