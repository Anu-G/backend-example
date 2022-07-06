package repository

import "gorm.io/gorm"

type BasicRepositoryInterface interface {
	Create(model interface{}) error
	FindById(model interface{}) error
	Update(model interface{}, with map[string]interface{}) error
	Delete(model interface{}) error
}

type basicRepository struct {
	db *gorm.DB
}

func NewBasicRepo(db *gorm.DB) BasicRepositoryInterface {
	return &basicRepository{
		db: db,
	}
}

func (bsr *basicRepository) Create(model interface{}) error {
	return bsr.db.Create(model).Error
}

func (bsr *basicRepository) FindById(model interface{}) error {
	return bsr.db.First(&model).Error
}

func (bsr *basicRepository) Update(model interface{}, with map[string]interface{}) error {
	return bsr.db.Model(&model).Updates(with).Error
}

func (bsr *basicRepository) Delete(model interface{}) error {
	return bsr.db.Delete(model).Error
}
