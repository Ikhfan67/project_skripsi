package komen

type KomenModel struct {
	Name    string 	`json:"name"`
	Email 	string 	`json:"email"`
	Phone	string	`json:"phone"`
	Message	string	`json:"message"`
}
type TransformedKomenModel struct {
	Id		int		`json:"id"`
	Name    string 	`json:"name"`
	Email 	string 	`json:"email"`
	Phone	string	`json:"phone"`
	Message	string	`json:"message"`
}