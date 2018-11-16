package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type (
	ProductModel struct {
		ID          uuid.UUID     `json:"id"`
		IDCategory  uuid.UUID     `json:"id_category"`
		Description string        `json:"description"`
		Price       int           `json:"price"`
		Stock       int           `json:"stock"`
		ProductName string        `json:"product_name"`
		CreatedAt   time.Time     `json:"created_at"`
		CreatedBy   uuid.UUID     `json:"created_by"`
		UpdateAt    pq.NullTime   `json:"update_at"`
		UpdateBy    uuid.NullUUID `json:"update_by"`
	}
)

func GetAllProduct(ctx context.Context, db *sql.DB) ([]ProductModel, error) {
	var productList []ProductModel
	query := `SELECT id, id_category, description, price, stock, product_name FROM "product"`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return productList, err
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductModel
		err := rows.Scan(
			&product.ID,
			&product.IDCategory,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.ProductName,
		)
		if err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}
	return productList, nil
}

func GetOneProduct(ctx context.Context, db *sql.DB, ID uuid.UUID) (ProductModel, error) {
	var product ProductModel
	query := `SELECT id, id_category, description, price, stock, product_name FROM "product" WHERE id=$1`
	err := db.QueryRowContext(ctx, query, ID).Scan(
		&product.ID,
		&product.IDCategory,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.ProductName,
	)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (product ProductModel) Insert(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	fmt.Println(product)

	query := `INSERT INTO "product"(id, id_category, description, price, stock, product_name, created_at, created_by)VALUES(uuid_generate_v4(), $1, $2, $3, $4, $5, now(), $6) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		product.IDCategory,
		product.Description,
		product.Price,
		product.Stock,
		product.ProductName,
		product.CreatedBy).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (product ProductModel) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE "product" SET(description, price, stock, product_name, update_at, update_by)=($1, $2, $3, $4, now(), $5) WHERE id=$6`
	_, err := db.ExecContext(ctx, query,
		product.Description,
		product.Price,
		product.Stock,
		product.ProductName,
		product.UpdateBy,
		product.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ctx context.Context, db *sql.DB, ID uuid.UUID) error {
	query := `DELETE FROM "product" WHERE id=$1`
	_, err := db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
