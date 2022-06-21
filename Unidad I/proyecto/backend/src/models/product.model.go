package models

import (
	"database/sql"
	"proyecto/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) GetProducts() (products []entities.Product, err error) {
	q := `SELECT
			p.ProductID,
			p.ProductName,
			s.CompanyName,
			c.CategoryName,
			p.UnitPrice,
			p.UnitsInStock
		FROM Products p
		INNER JOIN Suppliers s
		ON p.SupplierID = s.SupplierID
		INNER JOIN Categories c 
		ON p.CategoryID = c.CategoryID
		WHERE p.Discontinued = 0
	`
	rows, err := productModel.Db.Query(q)

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product

		for rows.Next() {
			var ProductID int
			var ProductName string
			var CompanyName string
			var CategoryName string
			var UnitPrice float64
			var UnitsInStock int

			err2 := rows.Scan(&ProductID, &ProductName, &CompanyName, &CategoryName, &UnitPrice, &UnitsInStock)

			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					ProductID:    ProductID,
					ProductName:  ProductName,
					CompanyName:  CompanyName,
					CategoryName: CategoryName,
					UnitPrice:    UnitPrice,
					UnitsInStock: UnitsInStock,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchProduct(keyword string) (products []entities.Product, err error) {
	rows, err := productModel.Db.Query("SELECT p.ProductID, p.ProductName, s.CompanyName, c.CategoryName, p.UnitPrice, p.UnitsInStock "+
		"FROM Products p "+
		"INNER JOIN Suppliers s "+
		"ON p.SupplierID = s.SupplierID "+
		"INNER JOIN Categories c "+
		"ON p.CategoryID = c.CategoryID "+
		"WHERE p.ProductName LIKE ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product

		for rows.Next() {
			var ProductID int
			var ProductName string
			var CompanyName string
			var CategoryName string
			var UnitPrice float64
			var UnitsInStock int

			err2 := rows.Scan(&ProductID, &ProductName, &CompanyName, &CategoryName, &UnitPrice, &UnitsInStock)

			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					ProductID:    ProductID,
					ProductName:  ProductName,
					CompanyName:  CompanyName,
					CategoryName: CategoryName,
					UnitPrice:    UnitPrice,
					UnitsInStock: UnitsInStock,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}
