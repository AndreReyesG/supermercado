// TODO: Caching templates.
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) showDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// TODO: 404 Not Found
	department, err := app.departments.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Department: department,
	}

	page := "show.html"
	tmpl, ok := app.deptTmplCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exits", page)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base", data)
}

func (app *application) listDepartments(w http.ResponseWriter, r *http.Request) {
	departments, err := app.departments.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Departments: departments,
	}

	page := "list.html"
	tmpl, ok := app.deptTmplCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exits", page)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base", data)
}

func (app *application) showCreateDepartment(w http.ResponseWriter, r *http.Request) {
	// GET /departments/new
	page := "new.html"
	tmpl, ok := app.deptTmplCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exits", page)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base", nil)
}

func (app *application) createDepartment(w http.ResponseWriter, r *http.Request) {
	// POST /departments
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("problem calling r.ParseForm, %s", err.Error()), http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	manager := r.PostForm.Get("manager")

	id, err := app.departments.Insert(name, manager)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem inserting product %s", err.Error()), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/departments/%d", id), http.StatusSeeOther)
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

func (app *application) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	// POST /departments/{id}/delete
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	err = app.departments.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}
