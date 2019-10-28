package service

import (
	"github.com/jinzhu/gorm"
	"github.com/koloo91/model"
	"time"
)

func CreateAccount(db *gorm.DB, account model.Account) (model.Account, error) {
	if err := db.Create(&account).Error; err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func GetAccounts(db *gorm.DB) ([]model.Account, error) {
	accounts := make([]model.Account, 0)
	if err := db.Order("name asc").Find(&accounts).Error; err != nil {
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

func UpdateCategory(db *gorm.DB, id string, category model.Category) (model.Category, error) {
	var existingCategory model.Category
	if err := db.Where("id = ?", id).First(&existingCategory).Error; err != nil {
		return model.Category{}, err
	}

	existingCategory.Name = category.Name
	existingCategory.Updated = time.Now()

	if err := db.Save(&existingCategory).Error; err != nil {
		return model.Category{}, err
	}

	return existingCategory, nil
}

func GetCategories(db *gorm.DB) ([]model.Category, error) {
	categories := make([]model.Category, 0)
	if err := db.Order("name asc").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateBooking(db *gorm.DB, booking model.Booking) (model.Booking, error) {
	if err := db.Create(&booking).Error; err != nil {
		return model.Booking{}, err
	}
	return booking, nil
}

func GetBookings(db *gorm.DB) ([]model.Booking, error) {
	bookings := make([]model.Booking, 0)
	if err := db.Order("date desc").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
