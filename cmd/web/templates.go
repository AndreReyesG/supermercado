package main

import (
	"html/template"
	"path/filepath"

	"github.com/AndreReyesG/supermercado/internal/data"
)

type templateData struct {
	Product     data.Product
	Products    []data.Product
	Department  data.Department
	Departments []data.Department
}

// TODO: Make it work with with all html files in pages directory.
// NOTE: Maybe need a tree.
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/products/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set.
		tmpl, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		tmpl, err = tmpl.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template.
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}

// TODO: newDepartmentTemplateCache is the same as newTemplateCache. REFACTOR!!
func newDepartmentTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/departments/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set.
		tmpl, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		tmpl, err = tmpl.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template.
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
