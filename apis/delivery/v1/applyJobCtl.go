package v1

import (
	"errors"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
	"github.com/tuongnguyen1209/poly-career-back/variable"
)

type ApplyJobCtl struct {
	services *services.Service
}

func InitApplyJobCtl(services *services.Service) *ApplyJobCtl {
	return &ApplyJobCtl{
		services: services,
	}
}

func (ctl *ApplyJobCtl) CreateApplyJobWithOldCv(c echo.Context) error {

	jwtInfo := helper.GetJwtFromContext(c)

	applyJob := &models.ApplyJob{}

	if err := c.Bind(applyJob); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	err := ctl.services.StudentJobService.CreateJobWithOldCv(applyJob, jwtInfo.Id)

	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, nil)

}

func (ctl *ApplyJobCtl) CreateApplyJobWithNewCv(c echo.Context) error {
	jwtInfo := helper.GetJwtFromContext(c)

	applyJob := &dto.ApplyJobWithNewCv{}

	if err := c.Bind(applyJob); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	err := ctl.services.StudentJobService.CreateJobWithNewCv(applyJob, jwtInfo.Id)

	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, nil)
}

func (ctl *ApplyJobCtl) GetStudentCv(c echo.Context) error {
	jwtInfo := helper.GetJwtFromContext(c)

	studentCvs, err := ctl.services.StudentJobService.GetCvByIdStudent(jwtInfo.Id)

	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, studentCvs)

}
func (ctl *ApplyJobCtl) GetCvByEmployerId(c echo.Context) error {
	jwtInfo := helper.GetJwtFromContext(c)

	jobId, _ := strconv.Atoi(c.QueryParam("job_id"))
	status, errStatus := strconv.Atoi(c.QueryParam("status"))

	if errStatus != nil {
		status = -1
	}

	cvFilter := &models.CVFilter{
		JobId:  jobId,
		Status: status,
	}

	timeZone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	if dateFrom, err := time.Parse(variable.DateFormat, c.QueryParam("date_from")); err == nil {
		from := time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, timeZone)
		cvFilter.DateFrom = &from
	}

	if dateTo, err := time.Parse(variable.DateFormat, c.QueryParam("date_to")); err == nil {
		to := time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 23, 59, 59, 0, timeZone)

		cvFilter.DateTo = &to
	}

	if page, err := strconv.Atoi(c.QueryParam("page")); err != nil {
		cvFilter.Page = 1
	} else {
		cvFilter.Page = page

	}
	if limit, err := strconv.Atoi(c.QueryParam("limit")); err != nil {
		cvFilter.Limit = 10
	} else {
		cvFilter.Limit = limit
	}

	studentCvs, err, total := ctl.services.StudentJobService.GetCvByEmployerId(jwtInfo.Id, cvFilter)

	if err != nil {
		return response.Error(c, err)
	}

	meta := ctl.services.MetaService.GetMeta(models.Job{}, cvFilter.Page, cvFilter.Limit, total)

	return response.SuccessWithMeta(c, studentCvs, meta)

}

func (ctl *ApplyJobCtl) GetCvById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		})
	}

	applyJob, err := ctl.services.StudentJobService.GetCvByCvId(id)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, applyJob)
}

func (ctl *ApplyJobCtl) UpdateStatusApplyJob(c echo.Context) error {
	applyJobStatus := &dto.ApplyJobStatus{}

	if err := c.Bind(applyJobStatus); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := ctl.services.StudentJobService.UpdateStatusApplyJob(id, applyJobStatus.Status); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, "ok")
}
