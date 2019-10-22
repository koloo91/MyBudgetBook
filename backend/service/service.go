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

func GetAccounts(db *gorm.DB) ([]model.Account, error) {
	accounts := make([]model.Account, 0)
	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
