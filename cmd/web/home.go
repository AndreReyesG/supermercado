package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}
