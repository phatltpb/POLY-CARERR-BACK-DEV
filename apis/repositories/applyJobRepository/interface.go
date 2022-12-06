package applyJobRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type ApplyJobRepositoryInterface interface {
	CreateApplyJob(applyJob *models.ApplyJob) error
	CheckingStudentApplyJob(idPost, idStudent int) bool
	GetCvByIdEmployer(id int, cvFilter *models.CVFilter) ([]models.ApplyJob, error, int)
	GetApplyJobByCVId(id int) (*models.ApplyJob, error)
	UpdateApplyJob(id int, applyJob *models.ApplyJob) error
}
