package repository

import (
	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	CreateUser(uc *entity.UserCredential) error
	FindUser(uc *entity.UserCredential) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepositoryInterface {
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) CreateUser(uc *entity.UserCredential) error {
	return ar.db.Create(&uc).Error
}

func (ar *authRepository) FindUser(uc *entity.UserCredential) error {
	uc.Encode()
	return ar.db.First(&uc, "user_name = ?", uc.UserName).Error
}
