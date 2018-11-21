package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type (
	CategoryModel struct {
		ID           uuid.UUID     `json:"id"`
		CategoryName string        `json:"category_name"`
		CreatedAt    time.Time     `json:"created_at"`
		CreatedBy    uuid.UUID     `json:"created_by"`
		UpdateAt     pq.NullTime   `json:"update_at"`
		UpdateBy     uuid.NullUUID `json:"update_by"`
	}
)

func GetAllCategory(ctx context.Context, db *sql.DB) ([]CategoryModel, error) {
	var categoryList []CategoryModel
	query := `SELECT id, category_name FROM "category"`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return categoryList, err
	}
	defer rows.Close()

	for rows.Next() {
		var category CategoryModel
		err := rows.Scan(
			&category.ID,
			&category.CategoryName,
		)
		if err != nil {
			return categoryList, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}

func GetOneCategory(ctx context.Context, db *sql.DB, ID uuid.UUID) (CategoryModel, error) {
	var category CategoryModel
	query := `SELECT id, category_name FROM "category" WHERE id=$1`
	err := db.QueryRowContext(ctx, query, ID).Scan(
		&category.ID,
		&category.CategoryName,
	)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (category CategoryModel) Insert(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO "category"(category_name, created_at, created_by)VALUES($1, now(), $2) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		category.CategoryName,
		category.CreatedBy).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (category CategoryModel) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE "category" SET(category_name, update_at, update_by)=($1, now(), $2) WHERE id=$3`
	_, err := db.ExecContext(ctx, query,
		category.CategoryName,
		category.UpdateBy,
		category.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(ctx context.Context, db *sql.DB, ID uuid.UUID) error {
	query := `DELETE FROM "category" WHERE id=$1`
	_, err := db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
