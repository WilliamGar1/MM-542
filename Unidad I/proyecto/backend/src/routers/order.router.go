package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
