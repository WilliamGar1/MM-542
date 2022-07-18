package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"proyecto/src/entities"
	"proyecto/src/models"
)

type Res struct {
	Msj     string `json:"Msj"`
	Success int64  `json:"Success"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {

	orderModel := models.OrderModel{
		Db: db,
	}

	orders, err2 := orderModel.GetOrders()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(orders)
	}
}

func NewOrder(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	orderModel := models.OrderModel{
		Db: db,
	}

	var o entities.Order

	json.Unmarshal(reqBody, &o)

	order := entities.Order{
		OrderID:        o.OrderID,
		CustomerID:     o.CustomerID,
		EmployeeID:     o.EmployeeID,
		OrderDate:      o.OrderDate,
		RequiredDate:   o.RequiredDate,
		ShippedDate:    o.ShippedDate,
		ShipVia:        o.ShipVia,
		Freight:        o.Freight,
		ShipName:       o.ShipName,
		ShipAddress:    o.ShipAddress,
		ShipCity:       o.ShipCity,
		ShipRegion:     o.ShipRegion,
		ShipPostalCode: o.ShipPostalCode,
		ShipCountry:    o.ShipCountry,
	}

	rowsAffected, err2 := orderModel.NewOrder(&order)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("rowsAffected: ", rowsAffected)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var data Res

		data = Res{
			Msj:     "Orden ingresada con exito",
			Success: 1,
		}

		json.NewEncoder(w).Encode(data)
	}
}

func GetOrderID(w http.ResponseWriter, r *http.Request) {

	orderModel := models.OrderModel{
		Db: db,
	}

	id, err2 := orderModel.GetOrderID()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(id)
	}
}

func NewOrderDetail(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	orderModel := models.OrderModel{
		Db: db,
	}

	var oD entities.OrderDetail

	json.Unmarshal(reqBody, &oD)

	orderDetail := entities.OrderDetail{

		OrderID:   oD.OrderID,
		ProductID: oD.ProductID,
		UnitPrice: oD.UnitPrice,
		Quantity:  oD.Quantity,
		Discount:  oD.Discount,
	}

	rowsAffected, err2 := orderModel.NewOrderDetail(&orderDetail)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("rowsAffected: ", rowsAffected)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var data Res

		data = Res{
			Msj:     "Detalle ingresado con exito",
			Success: 1,
		}

		json.NewEncoder(w).Encode(data)
	}
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	orderModel := models.OrderModel{
		Db: db,
	}

	customers, err2 := orderModel.GetCustomers()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	orderModel := models.OrderModel{
		Db: db,
	}

	employees, err2 := orderModel.GetEmployees()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employees)
	}
}

func GetShippers(w http.ResponseWriter, r *http.Request) {

	orderModel := models.OrderModel{
		Db: db,
	}

	shippers, err2 := orderModel.GetShippers()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shippers)
	}
}
