package models

import (
	"log"

	"github.com/rodolfobandeira/go-web-studies/db"
)

// Product struct - Custom struct to hold Products information
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// RetrieveAllProducts - Return a list of Products
func RetrieveAllProducts() []Product {
	db := db.SqliteConnection()

	rows, err := db.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}
	products := []Product{}

	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Fatal(err)
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()

	return products
}

// LoadProduct - Return a Product
func LoadProduct(id int) Product {
	db := db.SqliteConnection()

	row, err := db.Query("SELECT id, name, description, price, quantity FROM products WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}

	for row.Next() {
		var name, description string
		var price float64
		var quantity int

		row.Scan(&id, &name, &description, &price, &quantity)

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()

	return product
}

// SaveProduct - Saves product on Sqlite3
func SaveProduct(name, description string, price float64, quantity int) {
	db := db.SqliteConnection()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO products (name, description, price, quantity) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(name, description, price, quantity)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
	defer stmt.Close()
}

// UpdateProduct - Saves product on Sqlite3
func UpdateProduct(name, description string, price float64, quantity, id int) {
	db := db.SqliteConnection()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("UPDATE products SET name=?, description=?, price=?, quantity=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(name, description, price, quantity, id)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
	defer stmt.Close()
}

// DeleteProduct - Removes an product from Database
func DeleteProduct(id int) {
	db := db.SqliteConnection()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?;")
	if err != nil {
		panic(err.Error)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error)
	}
	tx.Commit()

	defer stmt.Close()
	defer db.Close()
}
