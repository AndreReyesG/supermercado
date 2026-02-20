package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// NOTE: is it good that the data argument type is 'any'?
// or must be a pointer to templateData?
func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data any) {
	tmpl, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exits", page)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := &bytes.Buffer{}

	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem executing template, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}
