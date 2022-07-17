package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"proyecto/src/entities"
	"proyecto/src/models"
)

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
	}
}
