package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndreReyesG/supermercado/internal/data"
)

func (app *application) showProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	product, err := app.products.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, fmt.Sprintf("problem getting the product %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	data := templateData{
		Product: product,
	}

	app.render(w, r, http.StatusOK, "products/show.html", data)
}

func (app *application) listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.products.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem getting products %s", err.Error()), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Products: products,
	}

	app.render(w, r, http.StatusOK, "products/list.html", data)
}

func (app *application) showCreateProduct(w http.ResponseWriter, r *http.Request) {
	// GET /products/new
	app.render(w, r, http.StatusOK, "products/new.html", nil)
}

func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	// POST /products
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	supplier := r.PostForm.Get("supplier")

	id, err := app.products.Insert(name, supplier)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem inserting product %s", err.Error()), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/products/%d", id), http.StatusSeeOther)
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	// POST /products/{id}/delete
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	err = app.products.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func (app *application) productNotFound(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "products/not_found.html", nil)
}

func (app *application) productNoDept(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "products/no_dept.html", nil)
}

func (app *application) productNoInDept(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "products/no_in_dept.html", nil)
}
