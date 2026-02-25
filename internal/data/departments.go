package data

import (
	"database/sql"
	"errors"
)

type Department struct {
	ID       int64
	Name     string
	Manager  string
	Products []Product
}

type DepartmentModel struct {
	DB *sql.DB
}

func (d *DepartmentModel) Insert(name, manager string) (int64, error) {
	query := `
        INSERT INTO departments (name, manager) 
        VALUES (?, ?)`
	result, err := d.DB.Exec(query, name, manager)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *DepartmentModel) Get(id int64) (Department, error) {
	query := `
		SELECT id, name, manager FROM departments
    	WHERE id = ?`

	row := d.DB.QueryRow(query, id)

	var department Department

	err := row.Scan(&department.ID, &department.Name, &department.Manager)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Department{}, ErrRecordNotFound
		} else {
			return Department{}, err
		}
	}

	return department, nil
}

func (d *DepartmentModel) GetAll() ([]Department, error) {
	query := `SELECT id, name, manager FROM departments`
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []Department

	for rows.Next() {
		var d Department
		err = rows.Scan(&d.ID, &d.Name, &d.Manager)
		if err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

// TODO: Refactor
func (d *DepartmentModel) Delete(id int64) error {
	query := `DELETE FROM departments WHERE id = ?`
	_, err := d.DB.Exec(query, id)
	return err
}

func (d *DepartmentModel) GetDepartmentsWithProducts() ([]Department, error) {
	query := `SELECT 
    d.id AS department_id,
    d.name AS department_name,
    p.id AS product_id,
    p.name AS product_name
FROM departments d
LEFT JOIN products p 
    ON p.department_id = d.id
ORDER BY d.id`
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	departmentsMap := make(map[int64]*Department)

	for rows.Next() {
		var deptID int64
		var deptName string
		var productID sql.NullInt64
		var productName sql.NullString

		err := rows.Scan(&deptID, &deptName,
			&productID, &productName)
		if err != nil {
			return nil, err
		}

		if _, exists := departmentsMap[deptID]; !exists {
			departmentsMap[deptID] = &Department{
				ID:       deptID,
				Name:     deptName,
				Products: []Product{},
			}
		}

		if productID.Valid {
			departmentsMap[deptID].Products = append(
				departmentsMap[deptID].Products,
				Product{
					ID:   productID.Int64,
					Name: productName.String,
				},
			)
		}
	}

	var result []Department
	for _, dept := range departmentsMap {
		result = append(result, *dept)
	}

	return result, nil
}

func (d *DepartmentModel) GetDeptWithProducts(id int64) (Department, error) {
	query := `
        SELECT d.id, d.name, d.manager, p.id, p.name, p.supplier, p.price 
        FROM departments as d
		LEFT JOIN products as p
		ON d.id = p.department_id
		WHERE d.id=?`
	rows, err := d.DB.Query(query, id)
	if err != nil {
		return Department{}, nil
	}
	defer rows.Close()

	var department Department

	for rows.Next() {
		var d Department
		var p Product
		var productID *int64
		var productName *string
		var productSupplier *string
		var productPrice *float64

		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Manager,
			&productID,
			&productName,
			&productSupplier,
			&productPrice,
		)
		if err != nil {
			return Department{}, nil
		}

		department.ID = d.ID
		department.Name = d.Name
		department.Manager = d.Manager
		if productID != nil {
			p.ID = *productID
			p.Name = *productName
			p.Supplier = *productSupplier
			p.Price = productPrice
			p.DepartmentID = &d.ID
			department.Products = append(department.Products, p)
		}
	}

	if err = rows.Err(); err != nil {
		return Department{}, err
	}

	return department, nil
}
