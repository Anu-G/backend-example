package repository

import (
	"errors"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type TableRepositoryInterface interface {
	Create(t *entity.Table) error
	FindById(t *entity.Table) error
	FindAll(by map[string]interface{}) ([]entity.Table, error)
	Update(t *entity.Table) error
	Delete(t *entity.Table) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepo(db *gorm.DB) TableRepositoryInterface {
	return &tableRepository{
		db: db,
	}
}

func (tr *tableRepository) Create(t *entity.Table) error {
	return tr.db.Create(&t).Error
}

func (tr *tableRepository) FindById(t *entity.Table) error {
	return tr.db.First(&t).Error
}

func (tr *tableRepository) FindAll(by map[string]interface{}) ([]entity.Table, error) {
	var tables []entity.Table
	res := tr.db.Where(by).Find(&tables)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tables, nil
		} else {
			return tables, err
		}
	}
	return tables, nil
}

func (tr *tableRepository) Update(t *entity.Table) error {
	return tr.db.Model(&t).Updates(&t).Error
}

func (tr *tableRepository) Delete(t *entity.Table) error {
	return tr.db.Delete(&t).Error
}
