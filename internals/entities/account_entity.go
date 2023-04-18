package entities

import (
	"nito/api/internals/utils/crypto_util"

	"gorm.io/gorm"
)

type Account struct {
	Base
	Username string `gorm:"type:varchar(40);" json:"username"`
	Password string `gorm:"type:varchar(60);" json:"password"`
}

func (e *Account) BeforeSave(db *gorm.DB) (err error) {
	crypto_util.HashPassword(&e.Password)
	return nil
}
