// TODO: Caching templates.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) listProducts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/products/list.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}

func (app *application) showCreateProduct(w http.ResponseWriter, r *http.Request) {
	// GET /product/new
	fmt.Fprint(w, "Mostrar formulario para dar de alta algun producto")
}

func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	// POST /product
	fmt.Fprint(w, "Procesar formulario")
}

func (app *application) showAssignDepartment(w http.ResponseWriter, r *http.Request) {
	// GET /products/{id}/assign-department
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/products/assign_department.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}

func (app *application) assignDepartment(w http.ResponseWriter, r *http.Request) {
	// POST /products/{id}/assign-department
	fmt.Fprint(w, "Procesar formulario para dar de alta un prod a un depto")
}

func (app *application) showAddPrice(w http.ResponseWriter, r *http.Request) {
	// GET /products/{id}/prices/new
	fmt.Fprint(w, "Formulario para asignar precio a un producto")
}

func (app *application) addPrice(w http.ResponseWriter, r *http.Request) {
	// POST /products/{id}/prices
	fmt.Fprint(w, "Procesar formulario para asignar un precio a un prod")
}
