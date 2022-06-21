package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proyecto/src/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	producModel := models.ProductModel{
		Db: db,
	}

	products, err2 := producModel.GetProducts()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func FilterProducts(w http.ResponseWriter, r *http.Request) {

	producModel := models.ProductModel{
		Db: db,
	}

	products, err2 := producModel.SearchProduct("g")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}
