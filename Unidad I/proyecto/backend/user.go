/*package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn = "sqlserver://sa:Admin123@localhost:1433?database=NORTHWND"

type Products struct {
	ProductID       int     `json:"ProductID"`
	ProductName     string  `json:"ProductName"`
	SupplierID      int     `json:"SupplierID"`
	CategoryID      int     `json:"CategoryID"`
	QuantityPerUnit string  `json:"QuantityPerU"`
	UnitPrice       float64 `json:"UnitPrice"`
	UnitsInStock    int     `json:"UnitsInStock"`
	UnitsOnOrder    int     `json:"UnitsOnOrder"`
	ReorderLevel    int     `json:"ReorderLevel"`
	Discontinued    int     `json:"Discontinued"`
}

func initMigration() {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to db")
	}

	db.AutoMigrate(&Products{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var products []Products

	db.Find(&products)

	json.NewEncoder(w).Encode(products)

}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

	var product Products

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a Valid Task")
	}

	json.Unmarshal(reqBody, &product)

	w.Header().Set("Content-Type", "application/json")

	//json.NewDecoder(r.Body).Decode(&product)

	db.Create(&product)

	json.NewEncoder(w).Encode(product)

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func delUser(w http.ResponseWriter, r *http.Request) {

}
*/