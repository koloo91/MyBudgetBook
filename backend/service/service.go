package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/koloo91/model"
	"github.com/koloo91/repository"
	"time"
)

const (
	weekly     = "WEEKLY"
	monthly    = "MONTHLY"
	quarterly  = "QUARTERLY"
	halfYearly = "HALF_YEARLY"
	yearly     = "YEARLY"

	UpdateStrategyOne = "ONE"
	UpdateStrategyAll = "ALL"

	DeleteStrategyOne = "ONE"
	DeleteStrategyAll = "ALL"
)

func CreateAccount(ctx context.Context, db *sql.DB, account model.Account) (model.Account, error) {
	if err := repository.InsertAccount(ctx, db, account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func GetAccounts(ctx context.Context, db *sql.DB) ([]model.Account, error) {
	return repository.QueryAccounts(ctx, db)
}

func CreateCategory(ctx context.Context, db *sql.DB, category model.Category) (model.Category, error) {
	if err := repository.InsertCategory(ctx, db, category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func UpdateCategory(ctx context.Context, db *sql.DB, id string, category model.Category) (model.Category, error) {
	if err := repository.UpdateCategory(ctx, db, id, category); err != nil {
		return model.Category{}, err
	}

	return repository.GetCategoryById(ctx, db, id)
}

func GetCategories(ctx context.Context, db *sql.DB) ([]model.Category, error) {
	return repository.QueryCategories(ctx, db)
}

func CreateBooking(ctx context.Context, db *sql.DB, booking model.Booking) (model.Booking, error) {
	tx, err := db.Begin()

	if err != nil {
		return model.Booking{}, err
	}

	if len(booking.StandingOrderPeriod) > 0 {
		years, months, days, err := yearsMonthsDaysToAdd(booking.StandingOrderPeriod)
		if err != nil {
			tx.Rollback()
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
			if err := repository.InsertBooking(ctx, tx, newBooking); err != nil {
				tx.Rollback()
				return model.Booking{}, err
			}
			newBooking.Date = newBooking.Date.AddDate(0, 0, -daysToAdd)
		}
	}

	if err := repository.InsertBooking(ctx, tx, booking); err != nil {
		tx.Rollback()
		return model.Booking{}, err
	}

	return booking, tx.Commit()
}

func UpdateBooking(ctx context.Context, db *sql.DB, id string, booking model.Booking, updateStrategy string) (model.Booking, error) {
	if updateStrategy == UpdateStrategyOne {
		return updateSingleBooking(ctx, db, id, booking)
	} else if updateStrategy == UpdateStrategyAll {
		return updateAllBookings(ctx, db, id, booking)
	} else {
		return model.Booking{}, fmt.Errorf("invalid updateStrategy '%s'", updateStrategy)
	}
}

func updateSingleBooking(ctx context.Context, db *sql.DB, id string, booking model.Booking) (model.Booking, error) {
	if err := repository.UpdateBooking(ctx, db, id, booking); err != nil {
		return model.Booking{}, err
	}
	return repository.GetBookingById(ctx, db, id)
}

func updateAllBookings(ctx context.Context, db *sql.DB, id string, booking model.Booking) (model.Booking, error) {

	existingBooking, err := repository.GetBookingById(ctx, db, id)
	if err != nil {
		return model.Booking{}, err
	}

	if err := repository.UpdateBookings(ctx, db, existingBooking.StandingOrderId, booking); err != nil {
		return model.Booking{}, err
	}

	return repository.GetBookingById(ctx, db, id)
}

func GetBookings(ctx context.Context, db *sql.DB, startDate time.Time, endDate time.Time) ([]model.Booking, error) {
	return repository.QueryBookings(ctx, db, startDate, endDate)
}

func DeleteBooking(ctx context.Context, db *sql.DB, id string, deleteStrategy string) error {
	if deleteStrategy == DeleteStrategyOne {
		return repository.DeleteBooking(ctx, db, id)
	} else if deleteStrategy == DeleteStrategyAll {
		return repository.DeleteBookings(ctx, db, id)
	} else {
		return fmt.Errorf("invalid deleteStrategy '%s'", deleteStrategy)
	}
}

func GetBalances(ctx context.Context, db *sql.DB) ([]model.AccountBalance, error) {
	return repository.QueryBalances(ctx, db, EndOfDay().UTC())
}

func GetMonthStatistics(ctx context.Context, db *sql.DB, year int) ([]model.MonthStatistic, error) {
	return repository.QueryMonthStatistics(ctx, db, BeginningOfYearWithYear(year), EndOfYearWithYear(year))
}

func GetCategoryStatistics(ctx context.Context, db *sql.DB, year int) ([]model.CategoryStatistic, error) {
	return repository.QueryCategoryStatistic(ctx, db, BeginningOfYearWithYear(year), EndOfYearWithYear(year))
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

func BeginningOfYearWithYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.Now().Location())
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

func EndOfYearWithYear(year int) time.Time {
	return BeginningOfYearWithYear(year).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
