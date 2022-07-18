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

func (orderModel OrderModel) GetOrderID() (int64, error) {

	row, err := orderModel.Db.Query("SELECT top 1 OrderID FROM Orders ORDER BY OrderID desc")

	if err != nil {
		return 0, err
	} else {
		type r struct {
			ID int64 `json:"id"`
		}

		var orderId r

		for row.Next() {
			var id int64

			err2 := row.Scan(&id)

			if err2 != nil {
				return 0, err2
			} else {
				orderId = r{
					ID: id,
				}
			}
		}
		return orderId.ID, nil
	}
}

func (orderModel OrderModel) NewOrderDetail(orderDetail *entities.OrderDetail) (int64, error) {

	result, err := orderModel.Db.Exec("INSERT INTO [Order Details] values(?,?,?,?,?)",
		orderDetail.OrderID, orderDetail.ProductID, orderDetail.UnitPrice,
		orderDetail.Quantity, orderDetail.Discount)

	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}

func (orderModel OrderModel) GetCustomers() (customers []entities.Customer, err error) {

	rows, err := orderModel.Db.Query("SELECT CustomerID, ContactName FROM Customers")

	if err != nil {
		return nil, err
	} else {
		var customers []entities.Customer

		for rows.Next() {
			var CustomerID string
			var ContactName string

			err2 := rows.Scan(&CustomerID, &ContactName)

			if err2 != nil {
				return nil, err2
			} else {
				customer := entities.Customer{
					CustomerID:  CustomerID,
					ContactName: ContactName,
				}
				customers = append(customers, customer)
			}
		}
		return customers, nil
	}
}

func (orderModel OrderModel) GetEmployees() (employees []entities.Employee, err error) {

	rows, err := orderModel.Db.Query("SELECT EmployeeID, CONCAT(FirstName, ' ', LastName) FullName FROM Employees")

	if err != nil {
		return nil, err
	} else {
		var employees []entities.Employee

		for rows.Next() {
			var EmployeeID int
			var FullName string

			err2 := rows.Scan(&EmployeeID, &FullName)

			if err2 != nil {
				return nil, err2
			} else {
				employee := entities.Employee{
					EmployeeID: EmployeeID,
					FullName:   FullName,
				}
				employees = append(employees, employee)
			}
		}
		return employees, nil
	}
}

func (orderModel OrderModel) GetShippers() (shippers []entities.Shipper, err error) {

	rows, err := orderModel.Db.Query("SELECT ShipperID, CompanyName From Shippers")

	if err != nil {
		return nil, err
	} else {
		var shippers []entities.Shipper

		for rows.Next() {
			var ShipperID int
			var CompanyName string

			err2 := rows.Scan(&ShipperID, &CompanyName)

			if err2 != nil {
				return nil, err2
			} else {
				shipper := entities.Shipper{
					ShipperID:   ShipperID,
					CompanyName: CompanyName,
				}
				shippers = append(shippers, shipper)
			}
		}
		return shippers, nil
	}
}
