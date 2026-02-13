package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", app.home)

	router.HandleFunc("GET /products", app.listProducts)
	router.HandleFunc("GET /products/{id}", app.showProduct)
	router.HandleFunc("GET /products/new", app.showCreateProduct)
	router.HandleFunc("POST /products", app.createProduct)
	router.HandleFunc("POST /products/{id}/delete", app.deleteProduct)

	router.HandleFunc("GET /departments", app.listDepartments)

	router.HandleFunc("GET /products/{id}/assign-department", app.showAssignDepartment)

	router.HandleFunc("GET /departments/{id}/prices", app.showDepartmentPrices)

	return router
}
