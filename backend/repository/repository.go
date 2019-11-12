package repository

import (
	"context"
	"database/sql"
	"github.com/koloo91/model"
	"time"
)

func InsertAccount(ctx context.Context, db *sql.DB, account model.Account) error {
	_, err := db.ExecContext(ctx, "INSERT INTO accounts (id, name, starting_balance, created, updated) VALUES ($1, $2, $3, $4, $5)", account.Id, account.Name, account.StartingBalance, account.Created, account.Updated)
	if err != nil {
		return err
	}
	return nil
}

func QueryAccounts(ctx context.Context, db *sql.DB) ([]model.Account, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, starting_balance, created, updated FROM accounts ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]model.Account, 0)
	var id, name string
	var startingBalance float64
	var created, updated time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &name, &startingBalance, &created, &updated); err != nil {
			return nil, err
		}

		result = append(result, model.Account{
			Id:              id,
			Name:            name,
			StartingBalance: startingBalance,
			Created:         created,
			Updated:         updated,
		})
	}

	return result, nil
}

func InsertCategory(ctx context.Context, db *sql.DB, category model.Category) error {
	_, err := db.ExecContext(ctx, "INSERT INTO categories (id, name, created, updated) VALUES ($1, $2, $3, $4)", category.Id, category.Name, category.Created, category.Updated)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(ctx context.Context, db *sql.DB, id string, category model.Category) error {
	_, err := db.ExecContext(ctx, "UPDATE categories SET name = $1, updated = $2 WHERE id = $3", category.Name, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
