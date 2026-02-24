package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) showAssignDepartment(w http.ResponseWriter, r *http.Request) {
	// GET /products/assign-department
	app.render(w, r, http.StatusOK, "products/assign_department.html", nil)
}

func (app *application) addAssignDepartmentForm(w http.ResponseWriter, r *http.Request) {
	products, err := app.products.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem getting products %s", err.Error()), http.StatusInternalServerError)
		return
	}

	departments, err := app.departments.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Products:    products,
		Departments: departments,
	}

	app.render(w, r, http.StatusOK, "products/new_assign_department.html", data)
}

func (app *application) addAssignDepartment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}

	departmentID, err := strconv.ParseInt(r.PostForm.Get("department_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}
	productID, _ := strconv.ParseInt(r.PostForm.Get("product_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = app.products.AssignDepartment(productID, departmentID)
	if err != nil {
		http.Error(w, "Unable to assing department", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func (app *application) deleteAssignDepartmentForm(w http.ResponseWriter, r *http.Request) {
	products, err := app.products.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem getting products %s", err.Error()), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Products: products,
	}
	app.render(w, r, http.StatusOK, "products/delete_product_from_department.html", data)
}

func (app *application) removeDepartment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}

	productID, _ := strconv.ParseInt(r.PostForm.Get("product_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = app.products.RemoveDepartment(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func (app *application) listProdsByDepartment(w http.ResponseWriter, r *http.Request) {
	departments, err := app.departments.GetDepartmentsWithProducts()
	if err != nil {
		http.Error(w, "Unable to fetch data", http.StatusInternalServerError)
		return
	}
	data := templateData{
		Departments: departments,
	}
	app.render(w, r, http.StatusOK, "departments/products.html", data)
}
