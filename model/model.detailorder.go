package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type (
	DetailOrderModel struct {
		ID        uuid.UUID     `json:"id"`
		IDOrder   uuid.UUID     `json:"id_order"`
		ProductID uuid.UUID     `json:"product_id"`
		Quantity  int           `json:"quantity"`
		SubTotal  int           `json:"sub_total"`
		CreatedAt time.Time     `json:"created_at"`
		CreatedBy uuid.UUID     `json:"created_by"`
		UpdateAt  pq.NullTime   `json:"update_at"`
		UpdateBy  uuid.NullUUID `json:"update_by"`
	}
)

func GetAllDetailOrder(ctx context.Context, db *sql.DB) ([]DetailOrderModel, error) {
	var detailOrderList []DetailOrderModel
	query := `SELECT id, id_order, product_id, quantity, sub_total FROM "detail_order"`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return detailOrderList, err
	}
	defer rows.Close()

	for rows.Next() {
		var detailOrder DetailOrderModel
		err := rows.Scan(
			&detailOrder.ID,
			&detailOrder.IDOrder,
			&detailOrder.ProductID,
			&detailOrder.Quantity,
			&detailOrder.SubTotal,
		)
		if err != nil {
			return detailOrderList, err
		}
		detailOrderList = append(detailOrderList, detailOrder)
	}
	return detailOrderList, nil
}

func GetOneDetailOrder(ctx context.Context, db *sql.DB, ID uuid.UUID) (DetailOrderModel, error) {
	var detailOrder DetailOrderModel
	query := `SELECT id, id_order, product_id, quantity, sub_total FROM "detail_order" WHERE id=$1`
	err := db.QueryRowContext(ctx, query, ID).Scan(
		&detailOrder.ID,
		&detailOrder.IDOrder,
		&detailOrder.ProductID,
		&detailOrder.Quantity,
		&detailOrder.SubTotal,
	)
	if err != nil {
		return detailOrder, err
	}
	return detailOrder, nil
}

func (detailOrder DetailOrderModel) Insert(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO "detail_order"(id_order, product_id, quantity, sub_total, created_at, created_by)VALUES($1, $2, $3, $4, now(), $5) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		detailOrder.IDOrder,
		detailOrder.ProductID,
		detailOrder.Quantity,
		detailOrder.SubTotal,
		detailOrder.CreatedBy).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (detailOrder DetailOrderModel) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE "detail_order" SET(quantity, sub_total, update_at, update_by)=($1, $2, now(), $3) WHERE id=$4`
	_, err := db.ExecContext(ctx, query,
		detailOrder.Quantity,
		detailOrder.SubTotal,
		detailOrder.UpdateBy,
		detailOrder.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDetailOrder(ctx context.Context, db *sql.DB, ID uuid.UUID) error {
	query := `DELETE FROM "detail_order" WHERE id=$1`
	_, err := db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
