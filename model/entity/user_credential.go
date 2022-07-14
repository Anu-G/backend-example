package entity

import (
	"encoding/base64"
	"log"

	"gorm.io/gorm"
)

type UserCredential struct {
	gorm.Model
	UserName     string `gorm:"size:50;unique;not null"`
	UserPassword string `gorm:"size:50;not null"`
	Email        string `gorm:"size:50;unique;not null"`
	CustomerID   uint
}

func (uc UserCredential) TableName() string {
	return "m_user_credential"
}

func (uc *UserCredential) Encode() {
	uc.UserPassword = base64.StdEncoding.EncodeToString([]byte(uc.UserPassword))
}

func (uc *UserCredential) Decode() {
	data, err := base64.StdEncoding.DecodeString(uc.UserPassword)
	if err != nil {
		log.Println(err)
	}

	uc.UserPassword = string(data)
}
