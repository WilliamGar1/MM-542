package models

import (
	"database/sql"
	"proyecto/src/entities"
)

type OrderModel struct {
	Db *sql.DB
}

func (orderModel OrderModel) GetOrders() (orders []entities.Order, err error) {

	rows, err := orderModel.Db.Query("SELECT * FROM Orders")

	if err != nil {
		return nil, err
	} else {
		var orders []entities.Order

		for rows.Next() {
			var OrderID int
			var CustomerID string
			var EmployeeID int
			var OrderDate string
			var RequiredDate string
			var ShippedDate any
			var ShipVia int
			var Freight float64
			var ShipName string
			var ShipAddress string
			var ShipCity string
			var ShipRegion any
			var ShipPostalCode any
			var ShipCountry string

			err2 := rows.Scan(&OrderID, &CustomerID, &EmployeeID, &OrderDate, &RequiredDate, &ShippedDate, &ShipVia, &Freight,
				&ShipName, &ShipAddress, &ShipCity, &ShipRegion, &ShipPostalCode, &ShipCountry)

			if err2 != nil {
				return nil, err2
			} else {
				order := entities.Order{
					OrderID:        OrderID,
					CustomerID:     CustomerID,
					EmployeeID:     EmployeeID,
					OrderDate:      OrderDate,
					RequiredDate:   RequiredDate,
					ShippedDate:    ShippedDate,
					ShipVia:        ShipVia,
					Freight:        Freight,
					ShipName:       ShipName,
					ShipAddress:    ShipAddress,
					ShipCity:       ShipCity,
					ShipRegion:     ShipRegion,
					ShipPostalCode: ShipPostalCode,
					ShipCountry:    ShipCountry,
				}
				orders = append(orders, order)
			}
		}
		return orders, nil
	}
}
