package repository

import (
	"errors"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type DiscountRepositoryInterface interface {
	Create(d *entity.Discount) error
	FindById(d *entity.Discount) error
	FindByIdPreload(c *entity.Discount, preload string) error
	FindAll(by map[string]interface{}) ([]entity.Discount, error)
	Update(d *entity.Discount) error
	Delete(d *entity.Discount) error
	DeleteAssociation(c *entity.Discount, table string, delVal interface{}) error
}

type discountRepository struct {
	db *gorm.DB
}

func NewDiscountRepo(db *gorm.DB) DiscountRepositoryInterface {
	return &discountRepository{
		db: db,
	}
}

func (dr *discountRepository) Create(d *entity.Discount) error {
	return dr.db.Create(&d).Error
}

func (dr *discountRepository) FindById(d *entity.Discount) error {
	return dr.db.First(&d).Error
}

func (dr *discountRepository) FindByIdPreload(c *entity.Discount, preload string) error {
	return dr.db.Preload(preload).First(&c).Error
}

func (dr *discountRepository) FindAll(by map[string]interface{}) ([]entity.Discount, error) {
	var discounts []entity.Discount
	res := dr.db.Where(by).Find(&discounts)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discounts, nil
		} else {
			return discounts, err
		}
	}
	return discounts, nil
}

func (dr *discountRepository) Update(d *entity.Discount) error {
	return dr.db.Model(&d).Updates(&d).Error
}

func (dr *discountRepository) Delete(d *entity.Discount) error {
	return dr.db.Delete(&d).Error
}

func (dr *discountRepository) DeleteAssociation(c *entity.Discount, table string, delVal interface{}) error {
	return dr.db.Model(&c).Association(table).Delete(delVal)
}
