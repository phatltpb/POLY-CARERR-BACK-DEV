package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	"github.com/tuongnguyen1209/poly-career-back/config"
	"github.com/tuongnguyen1209/poly-career-back/pkg/client"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
	"github.com/tuongnguyen1209/poly-career-back/variable"
)

type ProfileController struct {
	Services services.Service
}

func InitProfileCtl(service *services.Service) *ProfileController {
	return &ProfileController{
		Services: *service,
	}
}

// // student
func (ctl *ProfileController) GetMyProfileStudent(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	student, err := ctl.Services.StudentService.GetProfile(jwtToken.Id)

	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, student)

}

func (ctl *ProfileController) UpdateProfileStudent(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	var students models.Students
	if err := c.Bind(&students); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	if err := ctl.Services.StudentService.UpdateStudentProfile(jwtToken.Id, &students); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}

func (ctl *ProfileController) UpdateStudentDetail(c echo.Context) error {

	jwtInfo := helper.GetJwtFromContext(c)

	studentProfile := &models.StudentProfile{}

	if err := c.Bind(studentProfile); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	if err := ctl.Services.StudentService.CreateOrUpdateStudentDetail(jwtInfo.Id, studentProfile); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, "ok")
}

func (ctl *ProfileController) CreateStudentEducation(c echo.Context) error {

	jwtInfo := helper.GetJwtFromContext(c)

	studentEducation := &models.StudentEducation{}

	if err := c.Bind(studentEducation); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}

	studentEducation.StudentId = jwtInfo.Id

	if err := ctl.Services.StudentService.CreateStudentEducation(studentEducation); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, "ok")
}

func (ctl *ProfileController) UpdateStudentEducation(c echo.Context) error {

	studentEducation := &models.StudentEducation{}

	if err := c.Bind(studentEducation); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		})
	}

	if err := ctl.Services.StudentService.UpdateStudentEducation(id, studentEducation); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, "ok")
}

func (ctl *ProfileController) DeleteStudentEducation(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.IdIsNotExits),
		})
	}

	if err := ctl.Services.StudentService.DeleteStudentEducation(id); err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, "ok")
}

func (ctl *ProfileController) UpdateStudentEmail(c echo.Context) error {

	newEmail := &dto.EditEmailDto{}
	if err := c.Bind(newEmail); err != nil {
		return response.Error(c, &response.CustomError{Error: err})
	}
	jwtInfo := helper.GetJwtFromContext(c)

	if err := ctl.Services.StudentService.ChangeStudentEmail(jwtInfo.Id, newEmail.NewEmail); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "Bạn đã yêu cầu đổi Email. Hãy kiểm tra email của bạn!")
}
func (ctl *ProfileController) VerifyChangeStudentEmail(c echo.Context) error {
	config := config.GetConfig()
	token := c.Param("token")

	err := ctl.Services.StudentService.VerifyChangeStudentEmail(token)

	if err != nil {

		return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s?change_mail=fail", config.Link.LinkRenderict))
	}

	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s?change_mail=success", config.Link.LinkRenderict))
}

// // user
func (ctl *ProfileController) GetMyProfileEmployer(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)

	user, err := ctl.Services.UserService.GetProfileById(jwtToken.Id)

	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, user)
}

func (ctl *ProfileController) UpdateMyProfileEmployer(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	var user models.User
	if err := c.Bind(&user); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	if err := ctl.Services.UserService.UpdateUserProfile(jwtToken.Id, &user); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}

func (ctl *ProfileController) ChangePassword(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	var newPass dto.ChangePassword

	changeStr := fmt.Sprintf("%s-%d", variable.ChangePassword, jwtToken.Id)

	countStr, _ := client.GetValue(changeStr)
	count, _ := strconv.Atoi(countStr)

	if count == 5 {
		fmt.Println(">>>Log student: ", changeStr)
	}

	if err := c.Bind(&newPass); err != nil {
		client.SetValue(changeStr, count+1, time.Minute*5)
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	if err := ctl.Services.StudentService.ChangePassword(jwtToken.Id, &newPass); err != nil {
		client.SetValue(changeStr, count+1, time.Minute*5)
		return response.Error(c, err)
	}

	client.SetValue(changeStr, 0, time.Second)
	return response.Success(c, "ok")

}

func (ctl *ProfileController) ChangeAdminPassword(c echo.Context) error {
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	var newPass dto.ChangePassword
	changeStr := fmt.Sprintf("%s-%d", variable.ChangeAdminPassword, jwtToken.Id)

	countStr, _ := client.GetValue(changeStr)
	count, _ := strconv.Atoi(countStr)

	if count == 5 {
		fmt.Println(">>>Log student: ", changeStr)
	}

	if err := c.Bind(&newPass); err != nil {
		client.SetValue(changeStr, count+1, time.Minute*5)

		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	if err := ctl.Services.UserService.ChangePassword(jwtToken.Id, &newPass); err != nil {
		client.SetValue(changeStr, count+1, time.Minute*5)

		return response.Error(c, err)
	}
	client.SetValue(changeStr, 0, time.Second)

	return response.Success(c, "ok")

}
