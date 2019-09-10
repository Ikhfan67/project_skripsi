package order
import(
	"time"
)
type OrderModel struct {
	IdPesanan			int					`json:"idorder"`
	IdUser				int					`json:"iduser"`
	Tglorder   			time.Time			`json:"tglorder"`
	Status				string				`json:"status"`
	ImageOrder			string				`json:"image_order"`
	Total				int					`json:"total"`
	TglUpload			time.Time			`json:"tglupload"`
	Name				string				`json:"name"`
	DetailPesanan 		[]DetailPesanan		`json:"produk"`
}
type DetailPesanan struct{
	Id				int			`json:"id"`
	IdProduk 		int			`json:"id_produk"`
	IdPesanan		int			`json:"idpesanan"`
	Price			int			`json:"price"`
	Name			string		`json:"name"`
	Quantity		int			`json:"quantity"`
}
type SellsProduct struct{
	Id			int			`json:"id"`
	Tglorder   	string 		`json:"tglorder"`
	Jumlah 		int			`json:"jumlah"`
	Name		string		`json:"name"`
	Total		int			`json:"total"`
	Image		string		`json:"image"`
	Price		int			`json:"price"`
}
type ExSellsProduct struct{
	IdPesanan	int			`json:"id_pesanan"`
	Tglorder   	string 		`json:"tglorder"`
	Name	   	string 		`json:"name"`
	Quantity	string 		`json:"quantity"`
	Price		string 		`json:"price"`
	TotalPrice	string 		`json:"total_price"`
}
type DetailSells struct{
	// Id			int			`json:"id"`
	Name		string		`json:"name"`
	Jumlah 		int			`json:"jumlah"`
	Total		int			`json:"total"`
	Image		string		`json:"image"`
	Price		int			`json:"price"`
}
type TransformedOrder struct {
	IdPesanan	int					`json:"idorder"`
	IdUser		int					`json:"iduser"`
	Tglorder   	string 				`json:"tglorder"`
	Status		string				`json:"status"`
	ImageOrder	string				`json:"image_order"`
	Total		int					`json:"total"`
	TglUpload	string				`json:"tglupload"`
}