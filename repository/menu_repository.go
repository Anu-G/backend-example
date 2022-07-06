package repository

import (
	"errors"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type MenuRepositoryInterface interface {
	Create(m *entity.Menu) error
	CreatePrice(mp *entity.MenuPrice) error
	FindById(m *entity.Menu) error
	FindByIdPreload(m *entity.Menu, preload string) error
	FindByIdAndPrice(m *entity.Menu, mp *entity.MenuPrice, preload string) error
	FindLatestMenuPrice(m *entity.Menu) (entity.MenuPrice, error)
	FindMenuByMenuPrice(t *entity.MenuPrice) (entity.Menu, error)
	FindMenuPriceById(t *entity.MenuPrice) error
	FindAll(by map[string]interface{}) ([]entity.Menu, error)
	Update(m *entity.Menu) error
	Delete(m *entity.Menu) error
	DeleteAssociation(model *entity.Menu, table string, delVal interface{}) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) MenuRepositoryInterface {
	return &menuRepository{
		db: db,
	}
}

func (mr *menuRepository) Create(m *entity.Menu) error {
	return mr.db.Create(&m).Error
}

func (mr *menuRepository) CreatePrice(mp *entity.MenuPrice) error {
	return mr.db.Create(&mp).Error
}

func (mr *menuRepository) FindById(m *entity.Menu) error {
	return mr.db.First(&m).Error
}

func (mr *menuRepository) FindByIdPreload(m *entity.Menu, preload string) error {
	return mr.db.Preload(preload).First(&m).Error
}

func (mr *menuRepository) FindByIdAndPrice(m *entity.Menu, mp *entity.MenuPrice, preload string) error {
	if err := mr.db.First(&m).Error; err != nil {
		return err
	}

	return mr.db.Model(&m).Where(map[string]interface{}{"price": mp.Price}).Association(preload).Find(&mp)
}

func (mr *menuRepository) FindLatestMenuPrice(m *entity.Menu) (entity.MenuPrice, error) {
	var foundPrice entity.MenuPrice
	tx := mr.db.Model(&entity.Menu{}).Preload("MenuPrices").First(&m)
	for i, data := range m.MenuPrices {
		if i == 0 {
			foundPrice = data
		}
		if foundPrice.CreatedAt.Before(data.CreatedAt) {
			foundPrice = data
		}
	}

	return foundPrice, tx.Error
}

func (mr *menuRepository) FindMenuByMenuPrice(t *entity.MenuPrice) (entity.Menu, error) {
	var (
		err       error
		foundMenu entity.Menu
	)

	if err = mr.db.First(&t).Error; err != nil {
		return foundMenu, err
	}
	foundMenu.ID = t.MenuID

	if err = mr.db.First(&foundMenu).Error; err != nil {
		return foundMenu, err
	}
	return foundMenu, err
}

func (mr *menuRepository) FindMenuPriceById(t *entity.MenuPrice) error {
	return mr.db.First(&t).Error
}

func (mr *menuRepository) FindAll(by map[string]interface{}) ([]entity.Menu, error) {
	var menus []entity.Menu
	res := mr.db.Where(by).Find(&menus)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menus, nil
		} else {
			return menus, err
		}
	}
	return menus, nil
}

func (mr *menuRepository) Update(m *entity.Menu) error {
	return mr.db.Model(&m).Updates(&m).Error
}

func (mr *menuRepository) Delete(m *entity.Menu) error {
	return mr.db.Delete(&m).Error
}

func (mr *menuRepository) DeleteAssociation(m *entity.Menu, table string, delVal interface{}) error {
	return mr.db.Model(&m).Association(table).Delete(delVal)
}
