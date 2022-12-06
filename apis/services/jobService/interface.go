package jobService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type JobServiceInterface interface {
	GetAllJob(jobFilter *models.JobFilter) ([]models.Job, *response.CustomError, int)
	CreateJob(job *models.Job) (int, *response.CustomError)
	GetJobById(id int, studentId int) (*models.Job, *response.CustomError)
	UpdateJob(id int, job *models.Job, user *pkg.JwtToken) *response.CustomError
	GetJobFit(jobId, studentId, limit int) []models.Job
}
