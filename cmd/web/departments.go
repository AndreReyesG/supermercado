// TODO: Caching templates.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) listDepartments(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/departments/list.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}

func (app *application) showCreateDepartment(w http.ResponseWriter, r *http.Request) {
	// GET /departments/new
	fmt.Fprint(w, "Mostrar formulario para dar de alta un depto")
}

func (app *application) createDepartment(w http.ResponseWriter, r *http.Request) {
	// POST /departments
	fmt.Fprint(w, "Procesar formulario para dar de alta un depto")
}

func (app *application) showDepartmentPrices(w http.ResponseWriter, r *http.Request) {
	// GET /departments/{id}/prices
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/departments/prices.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}
