package entities

type Order struct {
	OrderID        int64   `json:"OrderID"`
	CustomerID     string  `json:"CustomerID"`
	EmployeeID     int     `json:"EmployeeID"`
	OrderDate      string  `json:"OrderDate"`
	RequiredDate   string  `json:"RequiredDate"`
	ShippedDate    any     `json:"ShippedDate"`
	ShipVia        int     `json:"ShipVia"`
	Freight        float64 `json:"Freight"`
	ShipName       string  `json:"ShipName"`
	ShipAddress    string  `json:"ShipAddress"`
	ShipCity       string  `json:"ShipCity"`
	ShipRegion     any     `json:"ShipRegion"`
	ShipPostalCode any     `json:"ShipPostalCode"`
	ShipCountry    string  `json:"ShipCountry"`
}

type OrderDetail struct {
	OrderID   int64   `json:"OrderID"`
	ProductID int64   `json:"ProductID"`
	UnitPrice float64 `json:"UnitPrice"`
	Quantity  int     `json:"Quantity"`
	Discount  float64 `json:"Discount"`
}

type Customer struct {
	CustomerID  string `json:"CustomerID"`
	ContactName string `json:"ContactName"`
}

type Employee struct {
	EmployeeID int    `json:"EmployeeID"`
	FullName   string `json:"FullName"`
}

type Shipper struct {
	ShipperID   int    `json:"ShipperID"`
	CompanyName string `json:"CompanyName"`
}
