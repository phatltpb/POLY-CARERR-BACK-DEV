package v1

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	"github.com/tuongnguyen1209/poly-career-back/pkg/myLog"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	"github.com/tuongnguyen1209/poly-career-back/pkg/role"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type JobCtl struct {
	services *services.Service
}

func InitJobCtl(services *services.Service) *JobCtl {
	return &JobCtl{
		services: services,
	}
}

func (ctl *JobCtl) GetAllJobs(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	companyId, _ := strconv.Atoi(c.QueryParam("company_id"))
	categoryId, _ := strconv.Atoi(c.QueryParam("category_id"))
	minSalary, errMin := strconv.Atoi(c.QueryParam("min"))
	maxSalary, errMax := strconv.Atoi(c.QueryParam("max"))
	level := c.QueryParam("level")
	experience := c.QueryParam("experience")
	position := c.QueryParam("position")

	search := c.QueryParam("search")
	if errMax != nil && errMin != nil {
		minSalary = -1
		maxSalary = -1
	}
	jobFilter := &models.JobFilter{
		CompanyId:  companyId,
		CategoryId: categoryId,
		Search:     search,
		Position:   position,
		Experience: experience,
		Level:      level,
		MinSalary:  minSalary,
		MaxSalary:  maxSalary,
		Pagination: struct {
			Page  int
			Limit int
		}{
			Page:  page,
			Limit: limit,
		},
	}

	jobs, err, total := ctl.services.JobService.GetAllJob(jobFilter)

	if err != nil {
		return response.Error(c, err)
	}

	meta := ctl.services.MetaService.GetMeta(models.Job{}, page, limit, total)

	return response.SuccessWithMeta(c, jobs, meta)
}

func (ctl *JobCtl) CreateJob(c echo.Context) error {
	job := &models.Job{}

	if err := c.Bind(job); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	if jwtToken.Role == role.Employer {
		job.UserId = jwtToken.Id
	}

	jobId, err := ctl.services.JobService.CreateJob(job)

	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, jobId)

}

func (ctl *JobCtl) GetJobById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		})
	}

	jwtInfo := helper.GetJwtFromContext(c)

	studentId, _ := strconv.Atoi(c.QueryParam("student_id"))

	myLog.Log(jwtInfo)
	myLog.Log(studentId)
	job, errGet := ctl.services.JobService.GetJobById(id, studentId)

	if errGet != nil {
		return response.Error(c, errGet)
	}

	return response.Success(c, job)
}

func (ctl *JobCtl) UpdateJob(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	job := &models.Job{}

	if err := c.Bind(job); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	if err := ctl.services.JobService.UpdateJob(id, job, jwtToken); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}

func (ctl *JobCtl) GetJobFit(c echo.Context) error {

	idJob, _ := strconv.Atoi(c.QueryParam("job_id"))
	idStudent, _ := strconv.Atoi(c.QueryParam("student_id"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if idJob <= 0 && idStudent <= 0 {
		return response.Success(c, []int{})
	}

	if limit == 0 {
		limit = 10
	}

	rs := ctl.services.JobService.GetJobFit(idJob, idStudent, limit)
	return response.Success(c, rs)
}
