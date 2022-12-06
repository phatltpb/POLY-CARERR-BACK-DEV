package studentWatchRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type StudentWatchRepoInterface interface {
	GetWatchByStudentId(studentId, jobId int) (*models.StudentWatch, error)
	CreateStudentWatch(sWatch *models.StudentWatch) error
	UpdateStudentWatch(id int, sWatch *models.StudentWatch) error
	CreateOrUpdateStudentWatch(sWatch *models.StudentWatch) error
	GetAllStudentWatch() ([]models.StudentWatch, error)
}
