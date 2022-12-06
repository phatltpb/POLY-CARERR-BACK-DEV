package studentCvRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type StudentCvRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *StudentCvRepository {
	return &StudentCvRepository{
		db: db,
	}
}

func (s *StudentCvRepository) CreateStudentCv(studentCv *models.StudentCV) (*models.StudentCV, error) {
	err := s.db.Create(studentCv).Error

	return studentCv, err
}

func (s *StudentCvRepository) GetCvByIdStudent(id int) ([]models.StudentCV, error) {

	studentCvs := []models.StudentCV{}

	err := s.db.Where("student_id = ?", id).Find(&studentCvs).Error

	return studentCvs, err
}
