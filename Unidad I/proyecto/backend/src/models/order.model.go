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
			var OrderID int64
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

func (orderModel OrderModel) NewOrder(order *entities.Order) (int64, error) {
	result, err := orderModel.Db.Exec("INSERT INTO Orders values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		order.CustomerID, order.EmployeeID, order.OrderDate, order.RequiredDate, order.ShippedDate,
		order.ShipVia, order.Freight, order.ShipName, order.ShipAddress, order.ShipCity, order.ShipRegion,
		order.ShipPostalCode, order.ShipCountry)

	if err != nil {
		return 0, err
	} else {
		order.OrderID, _ = result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}
