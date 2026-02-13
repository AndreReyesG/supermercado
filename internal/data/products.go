package data

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID       int64
	Name     string
	Supplier string
}

type ProductModel struct {
	DB *sql.DB
}

func (p *ProductModel) Insert(name, supplier string) (int, error) {
	query := `
        INSERT INTO products (name, supplier) 
        VALUES (?, ?)`

	result, err := p.DB.Exec(query, name, supplier)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *ProductModel) Get(id int) (Product, error) {
	query := `
		SELECT id, name, supplier FROM products
    	WHERE id = ?`

	row := p.DB.QueryRow(query, id)

	var product Product

	err := row.Scan(&product.ID, &product.Name, &product.Supplier)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Product{}, ErrRecordNotFound
		} else {
			return Product{}, err
		}
	}
	return product, nil
}

// TODO: Refactor
func (p *ProductModel) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := p.DB.Exec(query, id)
	return err
}

func (p *ProductModel) GetAll() ([]Product, error) {
	query := `SELECT id, name, supplier FROM products`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Supplier)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
