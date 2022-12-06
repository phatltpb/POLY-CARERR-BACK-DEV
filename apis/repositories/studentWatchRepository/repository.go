package studentWatchRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type StudentWatchRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *StudentWatchRepository {
	return &StudentWatchRepository{
		db: db,
	}
}

func (r *StudentWatchRepository) GetWatchByStudentId(studentId, jobId int) (*models.StudentWatch, error) {

	sWatch := &models.StudentWatch{}

	err := r.db.Where("student_id = ?", studentId).Where("job_id = ?", jobId).First(sWatch).Error

	return sWatch, err
}

func (r *StudentWatchRepository) CreateStudentWatch(sWatch *models.StudentWatch) error {
	sWatch.Count = 1
	return r.db.Create(sWatch).Error
}
func (r *StudentWatchRepository) UpdateStudentWatch(id int, sWatch *models.StudentWatch) error {
	return r.db.Table("").Where("id = ?", id).UpdateColumns(sWatch).Error
}
func (r *StudentWatchRepository) CreateOrUpdateStudentWatch(sWatch *models.StudentWatch) error {

	studentWatch, err := r.GetWatchByStudentId(sWatch.StudentId, sWatch.JobId)

	if err != nil {
		return r.CreateStudentWatch(sWatch)
	} else {
		newCount := studentWatch.Count + 1

		return r.UpdateStudentWatch(studentWatch.Id, &models.StudentWatch{
			Count: newCount,
		})
	}

}

func (r *StudentWatchRepository) GetAllStudentWatch() ([]models.StudentWatch, error) {
	studentWatch := []models.StudentWatch{}

	err := r.db.Order("student_id").Find(&studentWatch).Error
	return studentWatch, err
}
