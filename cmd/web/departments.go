package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndreReyesG/supermercado/internal/data"
)

func (app *application) showDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	department, err := app.departments.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, fmt.Sprintf("problem getting department %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	data := templateData{
		Department: department,
	}

	app.render(w, r, http.StatusOK, "departments/show.html", data)
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

	app.render(w, r, http.StatusOK, "departments/list.html", data)
}

func (app *application) showCreateDepartment(w http.ResponseWriter, r *http.Request) {
	// GET /departments/new
	app.render(w, r, http.StatusOK, "departments/new.html", nil)
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

func (app *application) departmentNotFound(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "departments/not_found.html", nil)
}
