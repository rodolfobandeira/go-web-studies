package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/rodolfobandeira/go-web-studies/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Index - Show all products
func Index(w http.ResponseWriter, r *http.Request) {
	products := models.RetrieveAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

// New - Add new Product
func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

// Edit - Edit a Product
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idOk, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert id or int", err)
	}

	product := models.LoadProduct(idOk)

	templates.ExecuteTemplate(w, "Edit", product)
}

// Update - Updates a Product
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		id := r.FormValue("id")

		priceOk, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Failed to convert price or float", err)
		}

		quantityOk, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Failed to convert quantity to integer", err)
		}

		idOk, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to convert id or int", err)
		}

		models.UpdateProduct(name, description, priceOk, quantityOk, idOk)
	}

	http.Redirect(w, r, "/", 302)
}

// Create - Create a new Product on DB
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceOk, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Failed to convert price or float", err)
		}

		quantityOk, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Failed to convert quantity to integer", err)
		}

		models.SaveProduct(name, description, priceOk, quantityOk)
	}

	http.Redirect(w, r, "/", 302)
}

// Delete - Removes an entry from the database. "ID" is required
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
	w.Header().Set("Expires", "0")

	id := r.URL.Query().Get("id")

	idOk, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert id to integer", err)
	}

	models.DeleteProduct(idOk)

	http.Redirect(w, r, "/", 302)
	// For SQLite3, autoincrements will create the same id again by default.
	// If we use 301 redirect, it will permanently cache and wont remove
	// products agan.
	// 302 is temporary redirect so it doesn't cache it working as expected
}
