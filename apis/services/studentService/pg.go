package studentService

import (
	"errors"
	"fmt"
	"time"

	"github.com/tuongnguyen1209/poly-career-back/apis/dto"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	repository "github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentRepository"
	"github.com/tuongnguyen1209/poly-career-back/config"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	mailer "github.com/tuongnguyen1209/poly-career-back/pkg/mailer"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type StudentService struct {
	StudentRepo repository.StudentRepositoryInterface
}

func Init(repositories *repositories.Repositories) *StudentService {
	return &StudentService{
		StudentRepo: repositories.StudentRepository,
	}
}

func (s *StudentService) GetProfile(id int) (*models.Students, *response.CustomError) {

	if id == 0 {
		return nil, &response.CustomError{
			Error: errors.New(validationMessage.UserIsNotExits),
			Code:  mystatus.BadRequest,
		}
	}

	user, err := s.StudentRepo.GetStudentById(id)

	if err != nil {
		return nil, &response.CustomError{
			Error: err,
			Code:  mystatus.BadRequest,
		}
	}

	return user, nil
}

func (s *StudentService) UpdateStudentProfile(id int, student *models.Students) *response.CustomError {
	err := s.StudentRepo.UpdateStudent(student, id)
	if err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.UserIsNotExits),
			Code:  mystatus.BadRequest,
		}
	}
	return nil
}

func (s *StudentService) CreateOrUpdateStudentDetail(idStudent int, detail *models.StudentProfile) *response.CustomError {

	studentDetail, err := s.StudentRepo.GetStudentDetailByStudentId(idStudent)

	if err != nil {
		detail.StudentId = idStudent
		if err := s.StudentRepo.CreateStudentDetail(detail); err != nil {
			return &response.CustomError{
				Error: err,
			}
		}
	} else {
		if err := s.StudentRepo.UpdateStudentDetail(studentDetail.Id, detail); err != nil {
			return &response.CustomError{
				Error: err,
			}
		}
	}

	return nil
}

func (s *StudentService) CreateStudentEducation(education *models.StudentEducation) *response.CustomError {

	err := s.StudentRepo.CreateStudentEducation(education)
	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}
	return nil
}
func (s *StudentService) UpdateStudentEducation(id int, education *models.StudentEducation) *response.CustomError {

	err := s.StudentRepo.UpdateStudentEducation(id, education)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}
	return nil
}

func (s *StudentService) DeleteStudentEducation(id int) *response.CustomError {

	err := s.StudentRepo.DeleteStudentEducation(id)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	return nil
}

func (s *StudentService) ChangeStudentEmail(idStudent int, newEmail string) *response.CustomError {
	config := config.GetConfig()

	if s.StudentRepo.CheckingEmail(newEmail) {
		return &response.CustomError{
			Error: errors.New(validationMessage.EmailIsExit),
		}
	}

	student, err := s.StudentRepo.GetStudentById(idStudent)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	newClaims := pkg.JwtTokenChangeEmail{
		Id:       idStudent,
		OldEmail: student.Email,
		NewEmail: newEmail,
	}

	jwt := pkg.JwtConfig{}

	token, err := jwt.Encode(newClaims, time.Minute*15)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	body := fmt.Sprintf("%s/api/v1/verify/change_email/%s", config.Link.HostName, token)

	mailer.MailChangeEmail(&mailer.DataMailer{
		Subject:   newEmail,
		Body:      body,
		Recipient: newEmail,
	})

	return nil
}

func (s *StudentService) VerifyChangeStudentEmail(token string) *response.CustomError {

	jwtEmailInfo := &pkg.JwtTokenChangeEmail{}

	jwt := &pkg.JwtConfig{}

	if _, err := jwt.DecodeWithInterface(token, jwtEmailInfo); err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	student, err := s.StudentRepo.GetStudentById(jwtEmailInfo.Id)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	if student.Email != jwtEmailInfo.OldEmail {
		return &response.CustomError{
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}

	if err := s.StudentRepo.UpdateStudent(&models.Students{
		Email: jwtEmailInfo.NewEmail,
	}, jwtEmailInfo.Id); err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	return nil
}

func (Us *StudentService) ChangePassword(id int, changePass *dto.ChangePassword) *response.CustomError {

	user, err := Us.StudentRepo.GetStudentById(id)

	if err != nil {
		return &response.CustomError{
			Error: err,
		}
	}

	if err := helper.CheckPasswordHash(user.Password, changePass.CurrentPassword); err != nil {
		return &response.CustomError{
			Error: errors.New(validationMessage.PasswordWrong),
		}
	}

	var newStudent models.Students
	newStudent.Password, _ = helper.Hash(changePass.NewPassword)
	if err := Us.StudentRepo.UpdateStudent(&newStudent, id); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return nil
}
