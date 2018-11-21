package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type (
	WishlistModel struct {
		ID         uuid.UUID     `json:"id"`
		CustomerID uuid.UUID     `json:"customer_id"`
		ProductID  uuid.UUID     `json:"product_id"`
		CreatedAt  time.Time     `json:"created_at"`
		CreatedBy  uuid.UUID     `json:"created_by"`
		UpdateAt   pq.NullTime   `json:"update_at"`
		UpdateBy   uuid.NullUUID `json:"update_by"`
	}
)

func GetAllWishlist(ctx context.Context, db *sql.DB) ([]WishlistModel, error) {
	var wishlistList []WishlistModel
	query := `SELECT id, customer_id, product_id FROM "wishlist"`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return wishlistList, err
	}
	defer rows.Close()

	for rows.Next() {
		var wishlist WishlistModel
		err := rows.Scan(
			&wishlist.ID,
			&wishlist.CustomerID,
			&wishlist.ProductID,
		)
		if err != nil {
			return wishlistList, err
		}
		wishlistList = append(wishlistList, wishlist)
	}
	return wishlistList, nil
}

func GetOneWishlist(ctx context.Context, db *sql.DB, ID uuid.UUID) (WishlistModel, error) {
	var wishlist WishlistModel
	query := `SELECT id, customer_id, product_id FROM "wishlist" WHERE id=$1`
	err := db.QueryRowContext(ctx, query, ID).Scan(
		&wishlist.ID,
		&wishlist.CustomerID,
		&wishlist.ProductID,
	)
	if err != nil {
		return wishlist, err
	}
	return wishlist, nil
}

func (wishlist WishlistModel) Insert(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO "wishlist"(customer_id, product_id, created_at, created_by)VALUES($1, $2, now(), $3) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		wishlist.CustomerID,
		wishlist.ProductID,
		wishlist.CreatedBy).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (wishlist WishlistModel) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE "wishlist" SET(customer_id, product_id, update_at, update_by)=($1, $2, now(), $3) WHERE id=$4`
	_, err := db.ExecContext(ctx, query,
		wishlist.CustomerID,
		wishlist.ProductID,
		wishlist.UpdateBy,
		wishlist.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteWishlist(ctx context.Context, db *sql.DB, ID uuid.UUID) error {
	query := `DELETE FROM "wishlist" WHERE id=$1`
	_, err := db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
