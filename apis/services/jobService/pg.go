package jobService

import (
	"errors"
	"time"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/companyRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/jobRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentWatchRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/userRepository"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	"github.com/tuongnguyen1209/poly-career-back/pkg/role"
	"github.com/tuongnguyen1209/poly-career-back/pkg/similarity"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type JobService struct {
	jobRepository          jobRepository.JobRepositoryInterface
	userRepository         userRepository.UserRepositoryInterface
	companyRepository      companyRepository.CompanyRepositoryInterface
	studentWatchRepository studentWatchRepository.StudentWatchRepoInterface
}

func Init(repositories *repositories.Repositories) *JobService {
	return &JobService{
		jobRepository:          repositories.JobRepository,
		userRepository:         repositories.UserRepository,
		companyRepository:      repositories.CompanyRepository,
		studentWatchRepository: repositories.StudentWatch,
	}
}

func (s *JobService) GetAllJob(jobFilter *models.JobFilter) ([]models.Job, *response.CustomError, int) {
	jobs, err, total := s.jobRepository.GetJobs(jobFilter)
	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}, total
	}
	return jobs, nil, total
}

func (s *JobService) CreateJob(job *models.Job) (int, *response.CustomError) {

	if !s.companyRepository.CheckingCompanyOfUserIsExits(job.UserId) {
		return 0, &response.CustomError{
			Error: errors.New(validationMessage.RequireCompany),
		}
	}
	validation := helper.InitValidation().Require("category_id", job.CategoryId).
		Require("user_id", job.UserId).Require("title", job.Title)

	if !validation.IsValid {
		return 0, &response.CustomError{
			Code:       mystatus.UnprocessableEntity,
			ErrorField: validation.ErrorField,
		}
	}

	job, err := s.jobRepository.CreateJob(job)

	if err != nil {
		return 0, &response.CustomError{
			Error: err,
		}
	}

	return job.Id, nil
}

func (s *JobService) GetJobById(id int, studentId int) (*models.Job, *response.CustomError) {

	if id == 0 {
		return nil, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}

	job, err := s.jobRepository.GetJobById(id)

	if err != nil {
		return nil, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	}
	if studentId > 0 {
		studentWatch := &models.StudentWatch{
			StudentId: studentId,
			JobId:     job.Id,
		}
		s.studentWatchRepository.CreateOrUpdateStudentWatch(studentWatch)
	}

	return job, nil
}

func (s *JobService) UpdateJob(id int, job *models.Job, user *pkg.JwtToken) *response.CustomError {

	if currentJob, err := s.jobRepository.GetJobById(id); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	} else if user.Role == role.Employer && user.Id != currentJob.UserId {
		return &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		}
	} else if time.Now().After(currentJob.CreatedAt.Add(time.Minute * 60 * 24 * 3)) {
		return &response.CustomError{
			Error: errors.New(validationMessage.TimeExpired),
		}
	}

	if err := s.jobRepository.UpdateJob(id, job); err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	return nil
}

func (s *JobService) GetJobFit(jobId, studentId, limit int) []models.Job {

	studentWatches, err := s.studentWatchRepository.GetAllStudentWatch()

	if err != nil {
		return []models.Job{}
	}

	listSim, err := similarity.GetSimilarity(studentWatches, jobId, studentId)

	if err != nil {
		return []models.Job{}
	}
	listId := []int{}
	for _, sim := range listSim {
		if len(listId) > limit {
			break
		}

		listId = append(listId, sim.Id)
	}

	listJobFit, err := s.jobRepository.GetJobFromListJobId(listId)

	if err != nil {
		return []models.Job{}
	}
	return listJobFit
}
