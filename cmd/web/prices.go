package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndreReyesG/supermercado/internal/data"
)

func (app *application) showDepartmentPrices(w http.ResponseWriter, r *http.Request) {
	// GET /prices
	app.render(w, r, http.StatusOK, "departments/prices.html", nil)
}

func (app *application) assignPriceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "departments/assign_price.html", nil)
}

func (app *application) assignPrice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}

	departmentID, _ := strconv.ParseInt(r.PostForm.Get("department_id"), 10, 64)
	productID, _ := strconv.ParseInt(r.PostForm.Get("product_id"), 10, 64)
	price, _ := strconv.ParseFloat(r.PostForm.Get("price"), 64)

	department, err := app.departments.Get(departmentID)
	if err == data.ErrRecordNotFound {
		http.Redirect(w, r, "/departments/not-found", http.StatusSeeOther)
		return
	}

	product, err := app.products.Get(productID)
	if err == data.ErrRecordNotFound {
		http.Redirect(w, r, "/products/not-found", http.StatusSeeOther)
		return
	}

	if product.DepartmentID == nil {
		http.Redirect(w, r, "/products/no-dept", http.StatusSeeOther)
		return
	}

	if *product.DepartmentID != department.ID {
		http.Redirect(w, r, "/products/no-in-dept", http.StatusSeeOther)
		return
	}

	err = app.products.AssignPrice(productID, price)
	if err != nil {
		http.Error(w, fmt.Sprintf("problema asignando precios %s", err.Error()), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/products/%d", productID), http.StatusSeeOther)
}

func (app *application) searchPriceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "departments/search_price.html", nil)
}

func (app *application) searchPrice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}

	departmentID, _ := strconv.ParseInt(r.PostForm.Get("department_id"), 10, 64)
	if _, err := app.departments.Get(departmentID); err == data.ErrRecordNotFound {
		http.Redirect(w, r, "/departments/not-found", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/departments/%d/prices", departmentID), http.StatusSeeOther)
}

func (app *application) listDeptPrices(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	department, err := app.departments.GetDeptWithProducts(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problemas llamando GetDeptWithProducts, %s", err.Error()), http.StatusBadRequest)
		return
	}

	data := templateData{
		Department: department,
	}

	app.render(w, r, http.StatusOK, "departments/list_prices.html", data)
}
