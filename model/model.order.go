package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type (
	OrderModel struct {
		ID         uuid.UUID     `json:"id"`
		CustomerID uuid.UUID     `json:"customer_id"`
		TotalPrice int           `json:"total_price"`
		CreatedAt  time.Time     `json:"created_at"`
		CreatedBy  uuid.UUID     `json:"created_by"`
		UpdateAt   pq.NullTime   `json:"update_at"`
		UpdateBy   uuid.NullUUID `json:"update_by"`
	}
)

func GetAllOrder(ctx context.Context, db *sql.DB) ([]OrderModel, error) {
	var orderList []OrderModel
	query := `SELECT id, customer_id, total_price FROM "orders"`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return orderList, err
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderModel
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.TotalPrice,
		)
		if err != nil {
			return orderList, err
		}
		orderList = append(orderList, order)
	}
	return orderList, nil
}

func GetOneOrder(ctx context.Context, db *sql.DB, ID uuid.UUID) (OrderModel, error) {
	var order OrderModel
	query := `SELECT id, customer_id, total_price FROM "orders" WHERE id=$1`
	err := db.QueryRowContext(ctx, query, ID).Scan(
		&order.ID,
		&order.CustomerID,
		&order.TotalPrice,
	)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (order OrderModel) Insert(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO "orders"(customer_id, total_price, created_at, created_by)VALUES($1, $2, now(), $3) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		order.CustomerID,
		order.TotalPrice,
		order.CreatedBy).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (order OrderModel) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE "orders" SET(total_price, update_at, update_by)=($1, now(), $2) WHERE id=$3`
	_, err := db.ExecContext(ctx, query,
		order.TotalPrice,
		order.UpdateBy,
		order.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(ctx context.Context, db *sql.DB, ID uuid.UUID) error {
	query := `DELETE FROM "orders" WHERE id=$1`
	_, err := db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
