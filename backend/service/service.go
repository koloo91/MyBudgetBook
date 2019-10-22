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

func CreateCategory(db *gorm.DB, category model.Category) (model.Category, error) {
	if err := db.Create(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func GetCategories(db *gorm.DB) ([]model.Category, error) {
	categories := make([]model.Category, 0)
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
