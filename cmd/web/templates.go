package main

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/AndreReyesG/supermercado/internal/data"
)

type templateData struct {
	Product     data.Product
	Products    []data.Product
	Department  data.Department
	Departments []data.Department
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/**/*.html")
	if err != nil {
		return nil, err
	}

	rootPages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, root := range rootPages {
		pages = append(pages, root)
	}

	for _, page := range pages {
		name := strings.TrimPrefix(page, "ui/html/pages/")

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
