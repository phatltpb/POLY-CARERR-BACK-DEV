package applyJobRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type ApplyJobRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *ApplyJobRepository {
	return &ApplyJobRepository{
		db: db,
	}
}

func (r *ApplyJobRepository) CreateApplyJob(applyJob *models.ApplyJob) error {

	return r.db.Create(applyJob).Error
}
func (r *ApplyJobRepository) CheckingStudentApplyJob(idPost, idStudent int) bool {

	var count int64

	err := r.db.Table("apply_jobs").Where("post_id = ?", idPost).
		Where("cv_id IN (?)", r.db.Table("student_cvs").Where("student_id = ?", idStudent).Select("id")).Count(&count).Error

	if err != nil {
		return false
	}

	return count != 0
}

func (s *ApplyJobRepository) GetCvByIdEmployer(id int, cvFilter *models.CVFilter) ([]models.ApplyJob, error, int) {

	applyJobCvs := []models.ApplyJob{}

	var count int64

	search := s.db.Model(&models.ApplyJob{}).Preload("StudentCV.Students.StudentProfile").
		Preload("Job", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "status", "created_at", "deleted_at")
		})

	if cvFilter.JobId != 0 {
		search.Where("post_id = ?", cvFilter.JobId)
	}

	if cvFilter.Status >= 0 {
		search.Where("status = ?", cvFilter.Status)
	}

	if cvFilter.DateFrom != nil {
		search.Where("date_apply >= ?", cvFilter.DateFrom)
	}

	if cvFilter.DateTo != nil {
		search.Where("date_apply <= ?", cvFilter.DateTo)
	}
	page := cvFilter.Page
	limit := cvFilter.Limit

	offset := (page - 1) * limit

	err := search.
		Where("post_id IN (?)",
			s.db.Model(&models.Job{}).Where("user_id = ?", id).Select("id")).
		Count(&count).
		Offset(offset).Limit(limit).Order("id desc").
		// Select("jobs.title").
		Find(&applyJobCvs).Error

	return applyJobCvs, err, int(count)
}

func (s *ApplyJobRepository) GetApplyJobByCVId(id int) (*models.ApplyJob, error) {

	applyCv := &models.ApplyJob{}

	err := s.db.Where("id = ?", id).
		Preload("StudentCV.Students.StudentProfile.Province").
		Preload("StudentCV.Students.StudentProfile.Category").
		Preload("StudentCV.Students.StudentEducation").
		Preload("StudentCV.Students.Location").
		Preload("Job").First(applyCv).Error

	return applyCv, err
}
func (s *ApplyJobRepository) UpdateApplyJob(id int, applyJob *models.ApplyJob) error {

	return s.db.Where("id = ?", id).Model(&models.ApplyJob{}).UpdateColumns(applyJob).Error

}
