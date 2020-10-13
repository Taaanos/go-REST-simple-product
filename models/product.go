package models

import (
	"database/sql"
)

// Product is the product containers that visitors view
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ByID makes a query with the given id
func (p *Product) ByID(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

// Create creates a product in the database
func (p *Product) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO products(name,price) VALUES($1,$2) RETURNING id", p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

// Update updates the product in the database
func (p *Product) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, p.ID)
	return err
}

// Delete deletes the product from the database
func (p *Product) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err
}

// ByRange returns a list of products with IDs in the range specified by start and count
// start indicates the starting point in the database and count the number of items to return
// from this particular point.
func ByRange(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
