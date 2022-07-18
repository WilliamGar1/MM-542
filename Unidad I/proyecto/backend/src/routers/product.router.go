package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"proyecto/src/entities"
	"proyecto/src/models"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)

	keyword := vars["l"]

	productModel := models.ProductModel{
		Db: db,
	}

	products, err3 := productModel.SearchProduct(keyword)

	if err3 != nil {
		fmt.Println(err3)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	productModel := models.ProductModel{
		Db: db,
	}

	var t entities.Product

	json.Unmarshal(reqBody, &t)

	product := entities.Product{
		ProductID:    t.ProductID,
		UnitPrice:    t.UnitPrice,
		UnitsInStock: t.UnitsInStock,
	}

	rowsAffected, err2 := productModel.UpdateProduct(&product)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("rowsAffected: ", rowsAffected)
	}
}

func UpdateProductStock(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	productModel := models.ProductModel{
		Db: db,
	}

	var t entities.Product

	json.Unmarshal(reqBody, &t)

	product := entities.Product{
		ProductID:    t.ProductID,
		UnitsInStock: t.UnitsInStock,
	}

	rowsAffected, err2 := productModel.UpdateProductStock(&product)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("rowsAffected: ", rowsAffected)
	}
}
