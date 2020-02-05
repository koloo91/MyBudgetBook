package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/koloo91/mybudgetbook/model"
	"github.com/koloo91/mybudgetbook/repository"
	"runtime"
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

func CreateAccount(ctx context.Context, db *sql.DB, userId string, account model.Account) (model.Account, error) {
	if err := repository.InsertAccount(ctx, db, userId, account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func GetAccounts(ctx context.Context, db *sql.DB, userId string) ([]model.Account, error) {
	return repository.QueryAccounts(ctx, db, userId)
}

func CreateCategory(ctx context.Context, db *sql.DB, userId string, category model.Category) (model.Category, error) {
	if err := repository.InsertCategory(ctx, db, userId, category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func UpdateCategory(ctx context.Context, db *sql.DB, userId string, id string, category model.Category) (model.Category, error) {
	if err := repository.UpdateCategory(ctx, db, userId, id, category); err != nil {
		return model.Category{}, err
	}

	return repository.GetCategoryById(ctx, db, userId, id)
}

func GetCategories(ctx context.Context, db *sql.DB, userId string) ([]model.Category, error) {
	return repository.QueryCategories(ctx, db, userId)
}

func CreateBooking(ctx context.Context, db *sql.DB, userId string, booking model.Booking) (model.Booking, error) {
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
			if err := repository.InsertBooking(ctx, tx, userId, newBooking); err != nil {
				tx.Rollback()
				return model.Booking{}, err
			}
			newBooking.Date = newBooking.Date.AddDate(0, 0, -daysToAdd)
		}
	}

	if err := repository.InsertBooking(ctx, tx, userId, booking); err != nil {
		tx.Rollback()
		return model.Booking{}, err
	}

	return booking, tx.Commit()
}

func UpdateBooking(ctx context.Context, db *sql.DB, userId string, id string, booking model.Booking, updateStrategy string) (model.Booking, error) {
	if updateStrategy == UpdateStrategyOne {
		return updateSingleBooking(ctx, db, userId, id, booking)
	} else if updateStrategy == UpdateStrategyAll {
		return updateAllBookings(ctx, db, userId, id, booking)
	} else {
		return model.Booking{}, fmt.Errorf("invalid updateStrategy '%s'", updateStrategy)
	}
}

func updateSingleBooking(ctx context.Context, db *sql.DB, userId string, id string, booking model.Booking) (model.Booking, error) {
	if err := repository.UpdateBooking(ctx, db, userId, id, booking); err != nil {
		return model.Booking{}, err
	}
	return repository.GetBookingById(ctx, db, userId, id)
}

func updateAllBookings(ctx context.Context, db *sql.DB, userId string, id string, booking model.Booking) (model.Booking, error) {

	existingBooking, err := repository.GetBookingById(ctx, db, userId, id)
	if err != nil {
		return model.Booking{}, err
	}

	if err := repository.UpdateBookings(ctx, db, userId, existingBooking.StandingOrderId, booking); err != nil {
		return model.Booking{}, err
	}

	return repository.GetBookingById(ctx, db, userId, id)
}

func GetBookings(ctx context.Context, db *sql.DB, userId string, startDate time.Time, endDate time.Time) ([]model.Booking, error) {
	return repository.QueryBookingsWithStartAndEndDate(ctx, db, userId, startDate, endDate)
}

func DeleteBooking(ctx context.Context, db *sql.DB, userId string, id string, deleteStrategy string) error {
	if deleteStrategy == DeleteStrategyOne {
		return repository.DeleteBooking(ctx, db, userId, id)
	} else if deleteStrategy == DeleteStrategyAll {
		return repository.DeleteBookings(ctx, db, userId, id)
	} else {
		return fmt.Errorf("invalid deleteStrategy '%s'", deleteStrategy)
	}
}

func GetBalances(ctx context.Context, db *sql.DB, userId string) ([]model.AccountBalance, error) {
	return repository.QueryBalances(ctx, db, userId, EndOfDay().UTC())
}

func GetMonthStatistics(ctx context.Context, db *sql.DB, userId string, year int) ([]model.MonthStatistic, error) {
	return repository.QueryMonthStatistics(ctx, db, userId, BeginningOfYearWithYear(year), EndOfYearWithYear(year))
}

func GetCategoryStatistics(ctx context.Context, db *sql.DB, userId string, year int) ([]model.CategoryStatistic, error) {
	return repository.QueryCategoryStatistic(ctx, db, userId, BeginningOfYearWithYear(year), EndOfYearWithYear(year))
}

func GetInboxEntriesWithPossibleMatches(ctx context.Context, db *sql.DB, userId string) ([]model.InboxEntry, error) {
	bookings, err := repository.QueryBookings(ctx, db, userId)
	if err != nil {
		return nil, err
	}

	inboxEntries, err := repository.QueryInboxEntries(ctx, db, userId)
	if err != nil {
		return nil, err
	}

	workerChannel := make(chan model.InboxEntry, runtime.NumCPU()*2)
	resultChannel := make(chan model.InboxEntry)

	go func() {
		for _, inboxEntry := range inboxEntries {
			workerChannel <- inboxEntry
		}
	}()

	go func() {
		for inboxEntry := range workerChannel {
			// TODO: start goroutine
			for _, booking := range bookings {
				// TODO: define matching
				// 1. wenn booking date oder value date == date +5
				//    bei +/- 1 Tag 							+3
				//    bei +/- 2 Tag 							+1
				//    sonst 									+0
				// 2. wenn amount == amount						+5
				// max 10 Punkte
				// Schwellwert: 80%
				// bei 0% anlegen beim Hauptkonto und Kategorie sonstiges
			}
		}
	}()

	results := make([]model.InboxEntry, 0, len(inboxEntries))
	select {
	case foo := <-resultChannel:
		results = append(results, foo)
		if len(results) == len(inboxEntries) {
			close(resultChannel)
			close(workerChannel)
		}
	case <-ctx.Done():
		close(resultChannel)
		close(workerChannel)
		return nil, errors.New("timeout")
	}

	return results, nil
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
