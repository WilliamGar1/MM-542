package entities

type Order struct {
	OrderID        int     `json:"OrderID"`
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
