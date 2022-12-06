package jobRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
)

type JobRepositoryInterface interface {
	GetJobs(jobFilter *models.JobFilter) ([]models.Job, error, int)
	CreateJob(job *models.Job) (*models.Job, error)
	GetJobById(id int) (*models.Job, error)
	UpdateJob(id int, job *models.Job) error
	GetJobFromListJobId(ids []int) ([]models.Job, error)
}
