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
	// TODO: start transaction
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
			if newBooking.Date.After(endDate) {
				break
			}

			daysToAdd := daysToAddForWeekday(newBooking.Date.Weekday())
			newBooking.Date = newBooking.Date.AddDate(0, 0, daysToAdd)
			if err := db.Create(&newBooking).Error; err != nil {
				return model.Booking{}, err
			}
			newBooking.Date = newBooking.Date.AddDate(0, 0, -daysToAdd)
		}
	}

	if err := db.Create(&booking).Error; err != nil {
		return model.Booking{}, err
	}

	return booking, nil
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

func GetBookings(db *gorm.DB, startDate time.Time, endDate time.Time) ([]model.Booking, error) {
	bookings := make([]model.Booking, 0)
	if err := db.Where("date BETWEEN ? AND ?", startDate, endDate).Order("date desc").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func GetBalances(db *gorm.DB) ([]model.AccountBalance, error) {
	rows, err := db.Raw("SELECT account_id, SUM(amount) + (SELECT starting_balance FROM accounts WHERE id = b.account_id) as balance FROM bookings b JOIN accounts on b.account_id = accounts.id WHERE date <= ? GROUP BY account_id;", EndOfDay().UTC()).Rows()
	if err != nil {
		return []model.AccountBalance{}, err
	}

	result := make([]model.AccountBalance, 0)
	var accountId string
	var balance float64

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&accountId, &balance); err != nil {
			return []model.AccountBalance{}, err
		}

		result = append(result, model.AccountBalance{
			AccountId: accountId,
			Balance:   balance,
		})
	}

	return result, nil
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

func daysToAddForWeekday(bookingWeekday time.Weekday) int {
	switch bookingWeekday {
	case time.Saturday:
		return 2
	case time.Sunday:
		return 1
	default:
		return 0
	}
}

func BeginningOfMonth() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
}

func BeginningOfQuarter() time.Time {
	month := BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func BeginningOfHalf() time.Time {
	month := BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

func BeginningOfYear() time.Time {
	y, _, _ := time.Now().Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, time.Now().Location())
}

func EndOfDay() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 59, time.Now().Location())
}

func EndOfMonth() time.Time {
	return BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func EndOfQuarter() time.Time {
	return BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func EndOfHalf() time.Time {
	return BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

func EndOfYear() time.Time {
	return BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}
