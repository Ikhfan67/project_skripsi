package products

type ProductModel struct {
	Name     	string 	`json:"name"`
	Description string 	`json:"description"`
	Price		int		`json:"price"`
	Image		string	`json:"image"`
}

type TransformedProduct struct {
	Id  		int    	`json:"id"`
	Name 		string 	`json:"name"`
	Description string  `json:"description"`
	Price 		int   	`json:"price"`
	Image		string	`json:"image"`
}