package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/koloo91/model"
	"time"
)

const (
	weekly     = "WEEKLY"
	monthly    = "MONTHLY"
	quarterly  = "QUARTERLY"
	halfYearly = "HALF_YEARLY"
	yearly     = "YEARLY"
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
	if len(booking.StandingOrderPeriod) > 0 {
		years, months, days, err := yearsMonthsDaysToAdd(booking.StandingOrderPeriod)
		if err != nil {
			return model.Booking{}, err
		}

		booking.StandingOrderId = uuid.New().String()
		endDate := time.Now().AddDate(5, 0, 0)

		newBooking := booking

		for {
			newBooking.Id = uuid.New().String()
			newBooking.Date = newBooking.Date.AddDate(years, months, days)
			// newBooking.Date.Weekday()
			if newBooking.Date.After(endDate) {
				break
			}

			if err := db.Create(&newBooking).Error; err != nil {
				return model.Booking{}, err
			}
		}
	}

	if err := db.Create(&booking).Error; err != nil {
		return model.Booking{}, err
	}

	return booking, nil
}

func yearsMonthsDaysToAdd(period string) (years int, months int, days int, err error) {
	years = 0
	months = 0
	days = 0
	err = nil

	switch period {
	case weekly:
		days = 7
	case monthly:
		months = 1
	case quarterly:
		months = 3
	case halfYearly:
		months = 6
	case yearly:
		years = 1
	default:
		err = fmt.Errorf("invalid order period '%s'", period)
	}

	return
}

func UpdateBooking(db *gorm.DB, id string, booking model.Booking) (model.Booking, error) {
	var existingBooking model.Booking
	if err := db.Where("id = ?", id).First(&existingBooking).Error; err != nil {
		return model.Booking{}, err
	}

	existingBooking.Title = booking.Title
	existingBooking.Comment = booking.Comment
	existingBooking.Amount = booking.Amount
	existingBooking.Date = booking.Date
	existingBooking.CategoryId = booking.CategoryId
	existingBooking.AccountId = booking.AccountId
	existingBooking.Updated = time.Now()

	if err := db.Save(&existingBooking).Error; err != nil {
		return model.Booking{}, err
	}

	return existingBooking, nil
}

func GetBookings(db *gorm.DB) ([]model.Booking, error) {
	bookings := make([]model.Booking, 0)
	if err := db.Order("date desc").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
