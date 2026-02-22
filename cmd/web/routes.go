package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", app.home)

	// Products
	router.HandleFunc("GET /products", app.listProducts)
	router.HandleFunc("GET /products/{id}", app.showProduct)
	router.HandleFunc("GET /products/new", app.showCreateProduct)
	router.HandleFunc("POST /products", app.createProduct)
	router.HandleFunc("POST /products/{id}/delete", app.deleteProduct)

	// Departments
	router.HandleFunc("GET /departments", app.listDepartments)
	router.HandleFunc("GET /departments/{id}", app.showDepartment)
	router.HandleFunc("GET /departments/new", app.showCreateDepartment)
	router.HandleFunc("POST /departments", app.createDepartment)
	router.HandleFunc("POST /departments/{id}/delete", app.deleteDepartment)

	// Assigments
	router.HandleFunc("GET /products/assign-department", app.showAssignDepartment)
	router.HandleFunc("GET /products/assign-department/new", app.addAssignDepartmentForm)
	router.HandleFunc("POST /products/assign-department", app.addAssignDepartment)
	router.HandleFunc("GET /products/assign-department/delete", app.deleteAssignDepartmentForm)
	router.HandleFunc("POST /products/assign-department/delete", app.removeDepartment)

	// Prices
	router.HandleFunc("GET /departments/{id}/prices", app.showDepartmentPrices)

	return router
}
