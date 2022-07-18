package models

import (
	"database/sql"
	"proyecto/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) GetProducts() (products []entities.Product, err error) {

	rows, err := productModel.Db.Query("SELECT * FROM vista_productos_descontinuados")

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
	rows, err := productModel.Db.Query("SELECT * FROM vista_productos_descontinuados "+
		"WHERE ProductName LIKE ?", "%"+keyword+"%")

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

func (productModel ProductModel) UpdateProduct(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("UPDATE Products SET UnitPrice = ?, UnitsInStock = ? WHERE ProductID = ?",
		product.UnitPrice, product.UnitsInStock, product.ProductID)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModel ProductModel) UpdateProductStock(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("UPDATE Products SET UnitsInStock = ? WHERE ProductID = ?",
		product.UnitsInStock, product.ProductID)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
