package repository

import (
	"errors"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (u StudentRepository) GetUserByEmail(email string) (*models.Students, error) {

	student := models.Students{}

	rs := u.db.Where("email=?", email).First(&student)

	return &student, rs.Error
}

func (u StudentRepository) CreateStudent(student *models.Students) (*models.Students, error) {

	rs := u.db.Create(student)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return student, nil
}

func (u StudentRepository) GetStudentById(id int) (*models.Students, error) {

	student := &models.Students{}

	rs := u.db.Preload("StudentProfile.Province").Preload("StudentEducation").
		Where("id=?", id).First(student)

	if rs.Error != nil {
		return nil, errors.New(validationMessage.UserIsNotExits)
	}

	return student, nil
}

func (u StudentRepository) CheckingEmail(email string) bool {

	var count int64
	if u.db.Model(&models.Students{}).Where("email=?", email).Count(&count).Error != nil {
		return false
	}
	return count > int64(0)
}
func (u StudentRepository) UpdateActive(id int) bool {
	if err := u.db.Where("id = ?", id).Updates(&models.Students{IsActive: true}).Error; err != nil {
		return false
	}
	return true
}

func (u StudentRepository) UpdateStudent(student *models.Students, id int) error {
	if err := u.db.Where("id = ?", id).UpdateColumns(student).Error; err != nil {
		return err
	}
	return nil
}

func (u *StudentRepository) GetStudentDetailByStudentId(studentId int) (*models.StudentProfile, error) {

	studentProfile := &models.StudentProfile{}

	err := u.db.Where("student_id = ?", studentId).First(studentProfile).Error

	return studentProfile, err
}
func (u *StudentRepository) CreateStudentDetail(studentProfile *models.StudentProfile) error {

	return u.db.Create(studentProfile).Error
}
func (u *StudentRepository) UpdateStudentDetail(id int, studentProfile *models.StudentProfile) error {

	return u.db.Where("id = ?", id).UpdateColumns(studentProfile).Error
}

func (u *StudentRepository) GetStudentEducationById(id int) (*models.StudentEducation, error) {
	education := &models.StudentEducation{}

	err := u.db.Where("id = ?", id).First(education).Error

	return education, err
}

func (u *StudentRepository) CreateStudentEducation(education *models.StudentEducation) error {
	return u.db.Create(education).Error
}

func (u *StudentRepository) UpdateStudentEducation(id int, education *models.StudentEducation) error {
	return u.db.Model(&models.StudentEducation{}).Where("id = ?", id).
		UpdateColumns(education).Error
}
func (u *StudentRepository) DeleteStudentEducation(id int) error {
	return u.db.Delete(&models.StudentEducation{}, id).Error
}
