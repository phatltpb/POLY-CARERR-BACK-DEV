package controller

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
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/mailer"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
	"github.com/tuongnguyen1209/poly-career-back/variable"
)

type AuthController struct {
	Services *services.Service
}

func Init(services *services.Service) *AuthController {
	return &AuthController{
		Services: services,
	}
}

func (ctl *AuthController) Login(c echo.Context) error {

	loginDto := dto.LoginDto{}

	ip := fmt.Sprintf("%s-%s", variable.LoginCount, c.RealIP())
	countStr, _ := client.GetValue(ip)
	count, _ := strconv.Atoi(countStr)

	if count == 5 {
		fmt.Println(">>>IP log: ", ip)
	}

	if count > 5 {

		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.LoginFalseManyTime),
		})
	}

	if err := c.Bind(&loginDto); err != nil {

		client.SetValue(ip, count+1, time.Minute*5)
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
		})
	}

	token, err := ctl.Services.AuthService.Login(&loginDto)

	if err != nil {
		client.SetValue(ip, count+1, time.Minute*5)
		return response.Error(c, err)
	}
	client.SetValue(ip, 0, time.Second)

	return response.Success(c, map[string]string{
		"token": token,
	})
}

func (ctl *AuthController) Register(c echo.Context) error {
	body := models.Students{}
	config := config.GetConfig()

	if err := c.Bind(&body); err != nil {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
		})
	}

	token, err := ctl.Services.AuthService.Register(&body)

	if err != nil {
		return response.Error(c, err)
	}
	mailer := fmt.Sprintf("%s/api/v1/auth/verify-account/%s", config.Link.HostName, token)
	mailInfo := &pkg.DataMailer{
		Subject:   body.FullName,
		Body:      mailer,
		Recipient: body.Email,
	}
	pkg.MailVerifyAccount(mailInfo)

	return response.Success(c, "Ok")
}
func (ctl *AuthController) VerifyRegister(c echo.Context) error {
	token := c.Param("token")
	config := config.GetConfig()
	if _, err := ctl.Services.AuthService.VerifyRegister(token); err != nil {
		c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/auth/login?status=Fail", config.Link.LinkRenderict))
		return nil
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/auth/login?status=Success", config.Link.LinkRenderict))
	return nil
}

// / forgot password student
func (ctl *AuthController) ForgotStudentPassword(c echo.Context) error {
	var email dto.Forgotpassword
	c.Bind(&email)
	config := config.GetConfig()

	token, err := ctl.Services.AuthService.ForgotStudentPassword(email.Email)
	if err != nil {
		return response.Error(c, err)
	}

	mailer := fmt.Sprintf("%s/auth/forgot/update_password/%s", config.Link.HostFontName, token)
	mailInfo := &pkg.DataMailer{
		Subject:   email.Email,
		Body:      mailer,
		Recipient: email.Email,
	}
	pkg.MailVerifyAccount(mailInfo)
	return response.Success(c, "Ok")
}

// reset password
func (ctl *AuthController) ResetPassword(c echo.Context) error {
	token := c.Param("token")
	var newPassword dto.ResetPassword
	if err := c.Bind(&newPassword); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
			Code:  mystatus.BadRequest,
		})
	}
	if err := ctl.Services.AuthService.ResetStudentPassword(token, newPassword.Password); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}

////////////////////////////////////////////////////////////////

func (ctl *AuthController) LoginAdmin(c echo.Context) error {

	loginDto := dto.LoginDto{}

	ip := fmt.Sprintf("%s-%s", variable.LoginCount, c.RealIP())
	countStr, _ := client.GetValue(ip)
	count, _ := strconv.Atoi(countStr)

	if count == 5 {
		fmt.Println(">>>IP log: ", ip)
	}
	if count > 5 {

		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.LoginFalseManyTime),
		})
	}

	if err := c.Bind(&loginDto); err != nil {
		client.SetValue(ip, count+1, time.Minute*5)
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
		})
	}

	token, err := ctl.Services.AuthService.LoginAdmin(&loginDto)

	if err != nil {
		client.SetValue(ip, count+1, time.Minute*5)
		return response.Error(c, err)
	}
	client.SetValue(ip, 0, time.Second)
	return response.Success(c, map[string]string{
		"token": token,
	})
}

func (ctl *AuthController) RegisterAdmin(c echo.Context) error {

	body := models.User{}
	config := config.GetConfig()
	if err := c.Bind(&body); err != nil {
		return response.Error(c, &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
		})
	}

	adminToken, err := ctl.Services.AuthService.RegisterAdmin(&body)
	if err != nil {
		return response.Error(c, err)
	}
	mailer := fmt.Sprintf("%s/api/v1/auth/admin/verify-account/%s", config.Link.HostName, adminToken)
	mailInfo := &pkg.DataMailer{
		Subject:   body.FullName,
		Body:      mailer,
		Recipient: body.Email,
	}
	pkg.MailVerifyAccount(mailInfo)
	return response.Success(c, "Ok")
}

func (ctl *AuthController) VerifyAdminRegister(c echo.Context) error {
	token := c.Param("token")
	config := config.GetConfig()
	if _, err := ctl.Services.AuthService.VerifyAdminRegister(token); err != nil {
		c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/auth/login?status=Fail", config.Link.LinkRenderict))
		return nil
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/auth/login/employer?status=Success", config.Link.LinkRenderict))
	return nil
}

// forgot password student
func (ctl *AuthController) ForgotAdminPassword(c echo.Context) error {
	var email dto.Forgotpassword
	c.Bind(&email)
	config := config.GetConfig()

	token, err := ctl.Services.AuthService.ForGotAdminPassword(email.Email)
	if err != nil {
		return response.Error(c, err)
	}
	mailer := fmt.Sprintf("%s/auth/employer/forgot/update_password/%s", config.Link.HostName, token)
	mailInfo := &pkg.DataMailer{
		Subject:   email.Email,
		Body:      mailer,
		Recipient: email.Email,
	}
	pkg.MailVerifyAccount(mailInfo)
	return response.Success(c, "Ok")
}

// reset password
func (ctl *AuthController) ResetAdminPassword(c echo.Context) error {
	token := c.Param("token")
	var newPassword dto.ResetPassword
	if err := c.Bind(&newPassword); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
			Code:  mystatus.BadRequest,
		})
	}
	if err := ctl.Services.AuthService.ResetAdminPassword(token, newPassword.Password); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}
