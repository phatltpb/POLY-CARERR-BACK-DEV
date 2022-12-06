package jobRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) GetJobs(jobFilter *models.JobFilter) ([]models.Job, error, int) {

	jobs := []models.Job{}

	page := jobFilter.Pagination.Page
	limit := jobFilter.Pagination.Limit

	offset := (page - 1) * limit

	search := r.db.Preload("Category").Preload("User.Company").Preload("Location")

	var total int64
	if jobFilter.CompanyId != 0 {
		search.Where("user_id IN (?)", r.db.Where("company_id = ?", jobFilter.CompanyId).Select("id").Find(&models.User{}))
	}

	if jobFilter.CategoryId != 0 {
		search.Where("category_id = ?", jobFilter.CategoryId)
	}

	switch {
	case jobFilter.MaxSalary == 0 && jobFilter.MinSalary == 0:
		search.Where("salary = ?", 0)
	case jobFilter.MinSalary > 0 && jobFilter.MaxSalary > 0:
		search.Where("salary >= ? AND salary <=?", jobFilter.MinSalary, jobFilter.MaxSalary)
	case jobFilter.MaxSalary == 0 && jobFilter.MinSalary != 0:
		search.Where("salary >= ?", jobFilter.MinSalary, jobFilter.MinSalary)
	}
	if jobFilter.Experience != "" {
		search.Where("experience like ? ", jobFilter.Experience)
	}
	if jobFilter.Level != "" {
		search.Where("level like ? ", jobFilter.Level)
	}
	if jobFilter.Position != "" {
		search.Where("position like ? ", jobFilter.Position)
	}

	if jobFilter.Search != "" {
		search.Where("category_id IN (?) OR id IN (?) OR user_id IN (?)",
			r.db.Select("id").Where("name like ?", "%"+jobFilter.Search+"%").Find(&models.Category{}),
			r.db.Select("id").Where("title like ?", "%"+jobFilter.Search+"%").Find(&models.Job{}),
			r.db.Select("id").Where("company_id IN (?)",
				r.db.Select("id").Where("name LIKE ?", "%"+jobFilter.Search+"%").
					Find(&models.Company{})).
				Find(&models.User{}))

	}

	if err := search.Model(&models.Job{}).
		Count(&total).
		Offset(offset).Limit(limit).Order("id desc").
		Find(&jobs).
		Error; err != nil {
		return nil, err, 0
	}

	return jobs, nil, int(total)
}

func SearchJobByNameCate(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			return db.Select("Id").Where("name LIKE {{.name}}").First(&models.Category{})
		}
		return db
	}
}
func (r *JobRepository) CreateJob(job *models.Job) (*models.Job, error) {

	if err := r.db.Create(job).Error; err != nil {
		return nil, err
	}
	return job, nil
}
func (r *JobRepository) GetJobById(id int) (*models.Job, error) {

	job := &models.Job{}

	err := r.db.Preload("Category").Preload("User.Company").Preload("Location").
		Where("id=?", id).First(job).Error
	return job, err
}

func (r *JobRepository) UpdateJob(id int, job *models.Job) error {

	err := r.db.Model(&models.Job{}).Where("id= ? ", id).Updates(job).Error

	return err
}

func (r *JobRepository) GetJobFromListJobId(ids []int) ([]models.Job, error) {

	jobs := []models.Job{}

	err := r.db.Preload("Category").Preload("User.Company").Preload("Location").
		Where("id IN (?)", ids).
		Find(&jobs).
		Error

	return jobs, err
}
