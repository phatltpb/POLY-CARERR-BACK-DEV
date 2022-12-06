package studentJobService

import (
	"errors"
	"time"

	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/applyJobRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/jobRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentCvRepository"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type StudentJobService struct {
	applyJobRepository applyJobRepository.ApplyJobRepositoryInterface
	studentCv          studentCvRepository.StudentCvRepositoryInterface
	jobRepository      jobRepository.JobRepositoryInterface
}

func Init(repositories *repositories.Repositories) *StudentJobService {
	return &StudentJobService{
		applyJobRepository: repositories.ApplyJobRepository,
		studentCv:          repositories.StudentCv,
		jobRepository:      repositories.JobRepository,
	}
}

func (r *StudentJobService) CreateJobWithOldCv(applyJob *models.ApplyJob, studentId int) *response.CustomError {
	if _, err := r.jobRepository.GetJobById(applyJob.PostId); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}
	if r.applyJobRepository.CheckingStudentApplyJob(applyJob.PostId, studentId) {
		return &response.CustomError{
			Error: errors.New(validationMessage.ApplyJobIsExit),
		}
	}

	applyJob.DateApply = time.Now()

	err := r.applyJobRepository.CreateApplyJob(applyJob)
	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}
	return nil
}

func (r *StudentJobService) CreateJobWithNewCv(applyJobDto *dto.ApplyJobWithNewCv, studentId int) *response.CustomError {

	if _, err := r.jobRepository.GetJobById(applyJobDto.PostId); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}

	if r.applyJobRepository.CheckingStudentApplyJob(applyJobDto.PostId, studentId) {
		return &response.CustomError{
			Error: errors.New(validationMessage.ApplyJobIsExit),
		}
	}

	studentCv := &models.StudentCV{
		StudentId: studentId,
		Title:     applyJobDto.FileName,
		Link:      applyJobDto.FileUrl,
	}

	studentCv, err := r.studentCv.CreateStudentCv(studentCv)
	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	err = r.applyJobRepository.CreateApplyJob(applyJobDto.ConvertToApplyJob(studentCv.Id))

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	return nil
}

func (r *StudentJobService) GetCvByIdStudent(id int) ([]models.StudentCV, *response.CustomError) {

	studentCvs, err := r.studentCv.GetCvByIdStudent(id)

	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}
	}

	return studentCvs, nil
}
func (r *StudentJobService) GetCvByEmployerId(id int, cvFilter *models.CVFilter) ([]models.ApplyJob, *response.CustomError, int) {

	applyJobCvs, err, count := r.applyJobRepository.GetCvByIdEmployer(id, cvFilter)

	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}, 0
	}

	return applyJobCvs, nil, count
}

func (r *StudentJobService) GetCvByCvId(id int) (*models.ApplyJob, *response.CustomError) {

	applyJob, err := r.applyJobRepository.GetApplyJobByCVId(id)
	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}
	}
	return applyJob, nil
}

func (r *StudentJobService) UpdateStatusApplyJob(id, status int) *response.CustomError {
	if _, err := r.applyJobRepository.GetApplyJobByCVId(id); err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	err := r.applyJobRepository.UpdateApplyJob(id, &models.ApplyJob{
		Status: status,
	})

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	return nil
}
