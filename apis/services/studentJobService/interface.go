package studentJobService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type StudentJobServiceInterface interface {
	CreateJobWithOldCv(applyJob *models.ApplyJob, studentId int) *response.CustomError
	CreateJobWithNewCv(applyJobDto *dto.ApplyJobWithNewCv, studentId int) *response.CustomError
	GetCvByIdStudent(id int) ([]models.StudentCV, *response.CustomError)
	GetCvByEmployerId(id int, cvFilter *models.CVFilter) ([]models.ApplyJob, *response.CustomError, int)
	GetCvByCvId(id int) (*models.ApplyJob, *response.CustomError)
	UpdateStatusApplyJob(id, status int) *response.CustomError
}
