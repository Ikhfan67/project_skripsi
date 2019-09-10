package auth

type Profil struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}
	
type TransformedProfil struct {
	Id  	 int    `json:"id"`
	Name	 string	`json:"name"`
	Phone	 string `json:"phone"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Email	 string	`json:"email"`
	Password string `json:"password"`
	Role	 string `json:"role"`
	Gambar	 string `json:"gambar"`
}
type TransformedProfila struct {
	Id  	 int    `json:"id_user"`
	Name	 string	`json:"name"`
	Phone	 string `json:"phone"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Email	 string	`json:"email"`
	Password string `json:"password"`
	Gambar	 string `json:"image"`
}