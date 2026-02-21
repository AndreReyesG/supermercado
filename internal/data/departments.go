package data

import (
	"database/sql"
	"errors"
)

type Department struct {
	ID      int64
	Name    string
	Manager string
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
