package repository

import (
	"errors"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type TrxTypeInterface interface {
	Create(tt *entity.TransactionType) error
	FindById(tt *entity.TransactionType) error
	FindAll(by map[string]interface{}) ([]entity.TransactionType, error)
	Update(tt *entity.TransactionType, with map[string]interface{}) error
}

type trxTypeRepository struct {
	db *gorm.DB
}

func NewTrxTypeRepo(db *gorm.DB) TrxTypeInterface {
	return &trxTypeRepository{
		db: db,
	}
}

func (ttr trxTypeRepository) Create(tt *entity.TransactionType) error {
	return ttr.db.Create(tt).Error
}

func (ttr *trxTypeRepository) FindById(tt *entity.TransactionType) error {
	return ttr.db.First(&tt).Error
}

func (ttr *trxTypeRepository) FindAll(by map[string]interface{}) ([]entity.TransactionType, error) {
	var trxTypes []entity.TransactionType
	res := ttr.db.Where(by).Find(&trxTypes)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return trxTypes, nil
		} else {
			return trxTypes, err
		}
	}
	return trxTypes, nil
}

func (ttr *trxTypeRepository) Update(tt *entity.TransactionType, with map[string]interface{}) error {
	return ttr.db.Model(&tt).Updates(with).Error
}
