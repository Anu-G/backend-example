package repository

import (
	"errors"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	Create(c *entity.Customer) error
	FindById(c *entity.Customer) error
	FindByPhone(c *entity.Customer) error
	FindByIdPreload(c *entity.Customer, preload string) error
	FindFirtstPreload(preload string, by map[string]interface{}) (entity.Customer, error)
	FindAll(by map[string]interface{}) ([]entity.Customer, error)
	Update(c *entity.Customer) error
	Delete(c *entity.Customer) error
	DeleteAssociation(c *entity.Customer, table string, delVal interface{}) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) CustomerRepositoryInterface {
	return &customerRepository{
		db: db,
	}
}

func (cr *customerRepository) Create(c *entity.Customer) error {
	return cr.db.Create(&c).Error
}

func (cr *customerRepository) FindById(c *entity.Customer) error {
	return cr.db.First(&c).Error
}

func (cr *customerRepository) FindByPhone(c *entity.Customer) error {
	return cr.db.First(&c, "mobile_phone_no = ?", c.MobilePhoneNo).Error
}

func (cr *customerRepository) FindByIdPreload(c *entity.Customer, preload string) error {
	return cr.db.Preload(preload).First(&c).Error
}

func (cr *customerRepository) FindFirtstPreload(preload string, by map[string]interface{}) (entity.Customer, error) {
	var customer entity.Customer
	res := cr.db.Preload(preload).Where(by).First(&customer)
	if err := res.Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (cr *customerRepository) FindAll(by map[string]interface{}) ([]entity.Customer, error) {
	var customers []entity.Customer
	res := cr.db.Where(by).Find(&customers)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customers, nil
		} else {
			return customers, err
		}
	}
	return customers, nil
}

func (cr *customerRepository) Update(c *entity.Customer) error {
	return cr.db.Model(&c).Updates(&c).Error
}

func (cr *customerRepository) Delete(c *entity.Customer) error {
	return cr.db.Delete(&c).Error
}

func (cr *customerRepository) DeleteAssociation(c *entity.Customer, table string, delVal interface{}) error {
	return cr.db.Model(&c).Association(table).Delete(delVal)
}
