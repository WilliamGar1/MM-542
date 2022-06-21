package entities

type Product struct {
	ProductID    int     `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	CompanyName  string  `json:"CompanyName"`
	CategoryName string  `json:"CategoryName"`
	UnitPrice    float64 `json:"UnitPrice"`
	UnitsInStock int     `json:"UnitsInStock"`
}

/*func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nprice: %0.1f\nname: %s", product.ProductID, product.UnitPrice, product.ProductName)
}*/
