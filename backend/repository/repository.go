package repository

import (
	"context"
	"database/sql"
	"github.com/koloo91/mybudgetbook/model"
	"time"
)

func InsertAccount(ctx context.Context, db *sql.DB, userId string, account model.Account) error {
	_, err := db.ExecContext(ctx, "INSERT INTO accounts (id, user_id, name, starting_balance, is_main, created, updated) VALUES ($1, $2, $3, $4, $5, $6, $7)", account.Id, userId, account.Name, account.StartingBalance, account.IsMain, account.Created, account.Updated)
	return err
}

func QueryAccounts(ctx context.Context, db *sql.DB, userId string) ([]model.Account, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, starting_balance, is_main, created, updated FROM accounts WHERE user_id = $1 ORDER BY name", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]model.Account, 0)
	var id, name string
	var startingBalance float64
	var isMain bool
	var created, updated time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &name, &startingBalance, &isMain, &created, &updated); err != nil {
			return nil, err
		}

		result = append(result, model.Account{
			Id:              id,
			Name:            name,
			StartingBalance: startingBalance,
			IsMain:          isMain,
			Created:         created,
			Updated:         updated,
		})
	}

	return result, nil
}

func InsertCategory(ctx context.Context, db *sql.DB, userId string, category model.Category) error {
	_, err := db.ExecContext(ctx, "INSERT INTO categories (id, user_id, name, created, updated) VALUES ($1, $2, $3, $4, $5)", category.Id, userId, category.Name, category.Created, category.Updated)
	return err
}

func UpdateCategory(ctx context.Context, db *sql.DB, userId string, id string, category model.Category) error {
	_, err := db.ExecContext(ctx, "UPDATE categories SET name = $1, updated = $2 WHERE user_id = $3 AND id = $4", category.Name, time.Now(), userId, id)
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryById(ctx context.Context, db *sql.DB, userId string, categoryId string) (model.Category, error) {
	row := db.QueryRowContext(ctx, "SELECT id, name, created, updated FROM categories WHERE user_id = $1 AND id = $2", userId, categoryId)
	var id, name string
	var created, updated time.Time
	if err := row.Scan(&id, &name, &created, &updated); err != nil {
		return model.Category{}, err
	}
	return model.Category{
		Id:      id,
		Name:    name,
		Created: created,
		Updated: updated,
	}, nil
}

func QueryCategories(ctx context.Context, db *sql.DB, userId string) ([]model.Category, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, created, updated FROM categories WHERE user_id = $1 ORDER BY name", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var id, name string
	var created, updated time.Time

	result := make([]model.Category, 0)
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created, &updated); err != nil {
			return nil, err
		}

		result = append(result, model.Category{
			Id:      id,
			Name:    name,
			Created: created,
			Updated: updated,
		})
	}

	return result, nil
}

func InsertBooking(ctx context.Context, db *sql.Tx, userId string, booking model.Booking) error {
	_, err := db.ExecContext(ctx, "INSERT INTO bookings(id, user_id, title, date, amount, category_id, account_id, standing_order_id, standing_order_period, standing_order_last_day, created, updated) VALUES ($1, $2, $3, $4, $5,$6, $7,$8, $9, $10, $11, $12)",
		booking.Id, userId, booking.Title, booking.Date, booking.Amount, booking.CategoryId, booking.AccountId, booking.StandingOrderId, booking.StandingOrderPeriod, booking.StandingOrderLastDay, booking.Created, booking.Updated)
	return err
}

func UpdateBooking(ctx context.Context, db *sql.DB, userId string, id string, booking model.Booking) error {
	_, err := db.ExecContext(ctx, "UPDATE bookings SET title = $1, date = $2, amount = $3, category_id = $4, account_id = $5, updated=now() WHERE user_id = $6 AND id =$7",
		booking.Title, booking.Date, booking.Amount, booking.CategoryId, booking.AccountId, userId, id)
	return err
}

func UpdateBookings(ctx context.Context, db *sql.DB, userId string, standingOrderId string, booking model.Booking) error {
	_, err := db.ExecContext(ctx, "UPDATE bookings SET title = $1, amount = $2, category_id = $3, account_id = $4, updated=now() WHERE user_id = $5 AND standing_order_id = $6",
		booking.Title, booking.Amount, booking.CategoryId, booking.AccountId, userId, standingOrderId)
	return err
}

func GetBookingById(ctx context.Context, db *sql.DB, userId string, bookingId string) (model.Booking, error) {
	row := db.QueryRowContext(ctx, "SELECT id, title, date, amount, category_id, account_id, standing_order_id, standing_order_period, standing_order_last_day, created, updated FROM bookings WHERE user_id = $1 AND id = $2", userId, bookingId)
	var id, title, categoryId, accountId, standingOrderId, standingOrderPeriod string
	var amount float64
	var date, standingOrderLastDay, created, updated time.Time
	if err := row.Scan(&id, &title, &date, &amount, &categoryId, &accountId, &standingOrderId, &standingOrderPeriod, &standingOrderLastDay, &created, &updated); err != nil {
		return model.Booking{}, err
	}
	return model.Booking{
		Id:                   id,
		Title:                title,
		Date:                 date,
		Amount:               amount,
		CategoryId:           categoryId,
		AccountId:            accountId,
		StandingOrderId:      standingOrderId,
		StandingOrderPeriod:  standingOrderPeriod,
		StandingOrderLastDay: standingOrderLastDay,
		Created:              created,
		Updated:              updated,
	}, nil
}

func QueryBookings(ctx context.Context, db *sql.DB, userId string) ([]model.Booking, error) {
	rows, err := db.QueryContext(ctx, `SELECT id,
												   title,
												   date,
												   amount,
												   category_id,
												   account_id,
												   standing_order_id,
												   standing_order_period,
												   standing_order_last_day,
												   created,
												   updated
											FROM bookings
											WHERE user_id = $1
											ORDER BY date DESC`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, title, categoryId, accountId, standingOrderId, standingOrderPeriod string
	var amount float64
	var date, standingOrderLastDay, created, updated time.Time

	result := make([]model.Booking, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &title, &date, &amount, &categoryId, &accountId, &standingOrderId, &standingOrderPeriod, &standingOrderLastDay, &created, &updated); err != nil {
			return nil, err
		}
		result = append(result, model.Booking{
			Id:                   id,
			Title:                title,
			Date:                 date,
			Amount:               amount,
			CategoryId:           categoryId,
			AccountId:            accountId,
			StandingOrderId:      standingOrderId,
			StandingOrderPeriod:  standingOrderPeriod,
			StandingOrderLastDay: standingOrderLastDay,
			Created:              created,
			Updated:              updated,
		})
	}

	return result, nil
}

func QueryBookingsWithStartAndEndDate(ctx context.Context, db *sql.DB, userId string, startDate time.Time, endDate time.Time) ([]model.Booking, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, title, date, amount, category_id, account_id, standing_order_id, standing_order_period, standing_order_last_day, created, updated FROM bookings WHERE user_id = $1 AND date BETWEEN $2 AND $3 ORDER BY date DESC", userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, title, categoryId, accountId, standingOrderId, standingOrderPeriod string
	var amount float64
	var date, standingOrderLastDay, created, updated time.Time

	result := make([]model.Booking, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &title, &date, &amount, &categoryId, &accountId, &standingOrderId, &standingOrderPeriod, &standingOrderLastDay, &created, &updated); err != nil {
			return nil, err
		}
		result = append(result, model.Booking{
			Id:                   id,
			Title:                title,
			Date:                 date,
			Amount:               amount,
			CategoryId:           categoryId,
			AccountId:            accountId,
			StandingOrderId:      standingOrderId,
			StandingOrderPeriod:  standingOrderPeriod,
			StandingOrderLastDay: standingOrderLastDay,
			Created:              created,
			Updated:              updated,
		})
	}

	return result, nil
}

func DeleteBooking(ctx context.Context, db *sql.DB, userId string, id string) error {
	_, err := db.ExecContext(ctx, "DELETE FROM bookings WHERE user_id = $1 AND id = $2", userId, id)
	return err
}

func DeleteBookings(ctx context.Context, db *sql.DB, userId string, id string) error {
	_, err := db.ExecContext(ctx, "DELETE FROM bookings WHERE standing_order_id = (SELECT standing_order_id FROM bookings WHERE user_id = $1 AND id = $2) AND date >= (SELECT date FROM bookings WHERE user_id = $1 AND id = $2)", userId, id)
	return err
}

func QueryBalances(ctx context.Context, db *sql.DB, userId string, endDate time.Time) ([]model.AccountBalance, error) {
	rows, err := db.QueryContext(ctx, "SELECT account_id, name, SUM(amount) + (SELECT starting_balance FROM accounts WHERE user_id = $1 AND id = b.account_id) as balance FROM bookings b JOIN accounts on b.account_id = accounts.id WHERE b.user_id = $1 AND date <= $2 GROUP BY account_id, name", userId, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accountId, name string
	var balance float64

	result := make([]model.AccountBalance, 0)
	for rows.Next() {
		if err := rows.Scan(&accountId, &name, &balance); err != nil {
			return nil, err
		}

		result = append(result, model.AccountBalance{
			AccountId: accountId,
			Name:      name,
			Balance:   balance,
		})
	}

	return result, nil
}

func QueryMonthStatistics(ctx context.Context, db *sql.DB, userId string, startDate, endDate time.Time) ([]model.MonthStatistic, error) {

	rows, err := db.QueryContext(ctx, `SELECT CASE
								   WHEN sum(negative_bookings.amount) IS NULL
									   THEN 0
								   ELSE sum(negative_bookings.amount)
								   END as negative_amount,
							   CASE
								   WHEN sum(positive_bookings.amount) IS NULL
									   THEN 0
								   ELSE sum(positive_bookings.amount)
								   END as positive_amount,
							   series.month
						FROM (SELECT 0 as amount, generate_series(1, 12) as month) as series
								 LEFT JOIN (SELECT sum(amount)                  as amount,
												   to_char(date, 'MM')::INTEGER as month,
												   extract(year from date)      as year
											FROM bookings
											WHERE user_id = $1 AND date BETWEEN $2 AND $3
											  AND amount < 0
											GROUP BY month, year) as negative_bookings
										   ON series.month = negative_bookings.month
								 LEFT JOIN (SELECT sum(amount)                  as amount,
												   to_char(date, 'MM')::INTEGER as month,
												   extract(year from date)      as year
											FROM bookings
											WHERE user_id = $1 AND date BETWEEN $2 AND $3
											  AND amount >= 0
											GROUP BY month, year) as positive_bookings
										   ON series.month = positive_bookings.month
						GROUP BY series.month
						ORDER BY series.month DESC;`, userId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses, incomes float64
	var month int

	result := make([]model.MonthStatistic, 0)
	for rows.Next() {

		if err := rows.Scan(&expenses, &incomes, &month); err != nil {
			return nil, err
		}

		result = append(result, model.MonthStatistic{
			Expenses: expenses,
			Incomes:  incomes,
			Month:    month,
		})
	}

	return result, nil
}

func QueryCategoryStatistic(ctx context.Context, db *sql.DB, userId string, startDate, endDate time.Time) ([]model.CategoryStatistic, error) {

	rows, err := db.QueryContext(ctx, `SELECT c.name, ABS(SUM(amount)) as category_sum
									FROM bookings
									JOIN categories c on bookings.category_id = c.id
									WHERE c.user_id = $1 AND date BETWEEN $2 AND $3
									AND account_id = (SELECT id FROM accounts WHERE accounts.user_id = $1 AND is_main = true)
									GROUP BY category_id, c.name
									ORDER BY category_sum DESC;`, userId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var name string
	var sum float64

	result := make([]model.CategoryStatistic, 0)

	for rows.Next() {
		if err := rows.Scan(&name, &sum); err != nil {
			return nil, err
		}

		result = append(result, model.CategoryStatistic{
			Name: name,
			Sum:  sum,
		})
	}

	return result, nil
}

func QueryInboxEntries(ctx context.Context, db *sql.DB, userId string) ([]model.InboxEntry, error) {
	rows, err := db.QueryContext(ctx, `SELECT id,
												   booking_date,
												   value_date,
												   intended_use,
												   amount,
												   created,
												   updated
											FROM inbox
											WHERE user_id = $1
											ORDER BY booking_date DESC`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, intendedUse string
	var amount float64
	var bookingDate, created, updated time.Time
	var valueDateNullable sql.NullTime

	result := make([]model.InboxEntry, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &bookingDate, &valueDateNullable, &intendedUse, &amount, &created, &updated); err != nil {
			return nil, err
		}

		var valueDate *time.Time
		if valueDateNullable.Valid {
			valueDate = &valueDateNullable.Time
		}

		result = append(result, model.InboxEntry{
			Id:          id,
			BookingDate: bookingDate,
			ValueDate:   valueDate,
			IntendedUse: intendedUse,
			Amount:      amount,
			Created:     created,
			Updated:     updated,
		})
	}

	return result, nil
}
