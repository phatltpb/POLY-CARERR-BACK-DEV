package categoryRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	categories := []models.Category{}

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CreateCate(cate *models.Category) (*models.Category, error) {
	if err := r.db.Create(cate).Error; err != nil {
		return nil, err
	}
	return cate, nil
}

func (r *CategoryRepository) GetById(id int) (*models.Category, error) {
	var cate *models.Category
	if err := r.db.Where("id = ?", id).First(&cate).Error; err != nil {
		return nil, err
	}
	return cate, nil
}

func (r *CategoryRepository) UpdateCate(id int, cate *models.Category) error {
	if err := r.db.Where("id = ?", id).UpdateColumns(cate).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCate(id int) error {
	if err := r.db.Delete(&models.Category{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
