package service

import (
	"github.com/jinzhu/gorm"
	"github.com/koloo91/model"
)

func CreateAccount(db *gorm.DB, account model.Account) (model.Account, error) {
	if err := db.Create(&account).Error; err != nil {
		return model.Account{}, err
	}
	return account, nil
}
