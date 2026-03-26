package config

import (
	"backend/models"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetProducts(db *sql.DB) ([]models.Product, error) {
	query := "SELECT id, name, price, ratings, image_url, is_available, total_sold FROM products"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Ratings,
			&product.ImageURL,
			&product.IsAvailable,
			&product.TotalSold,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetCategories(db *sql.DB) ([]models.Category, error) {
	query := "SELECT id, name, icon FROM categories ORDER BY name ASC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var category models.Category
		var icon sql.NullString

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&icon,
		)
		if err != nil {
			return nil, err
		}

		if icon.Valid {
			category.Icon = icon.String
		} else {
			category.Icon = ""
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
