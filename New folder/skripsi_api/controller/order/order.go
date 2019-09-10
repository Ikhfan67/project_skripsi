package controller

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"bytes"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin/binding"
	
	"github.com/360EntSecGroup-Skylar/excelize"

	conn "skripsi_api/config"
	m "skripsi_api/models/order"
	
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" 
)

func init() {
	conn.Connect()
}

func CreateOrder(c *gin.Context) {
	var buffer bytes.Buffer
	var Order m.OrderModel

	if err:= c.ShouldBindJSON(&Order); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var now = time.Now()
	
	// Insert Pesanan

	stmt, err := conn.DB.Prepare("insert into pesanan (tglorder, id_user, status) values (?,?,?)")
	
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := stmt.Exec(now, Order.IdUser, Order.Status)
	idPesanan, _ := res.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}
	
	
	// Insert Produk


	data := Order.DetailPesanan

	vals := []interface{}{}
	for _, row := range data {
		vals = append(vals, idPesanan, row.IdProduk, row.Quantity)
	}

	sqlStr := `INSERT INTO detail_pesanan(id_pesanan, id_produk, quantity) VALUES %s`
	sqlStr = ReplaceSQL(sqlStr, "(?,?,?)", len(data))

	//Prepare and execute the statement
	stmt, _ = conn.DB.Prepare(sqlStr)

	res, err = stmt.Exec(vals...)
	
	if err != nil {
		log.Panic(err)
	}

	id, _ := res.LastInsertId()
	
	buffer.WriteString(strconv.Itoa(int(id)))
	buffer.WriteString(" ")
	buffer.WriteString(Order.Status)
	defer stmt.Close()
	order := buffer.String()
	fmt.Println(order)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully add", order),
	})
}
func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
	return strings.TrimSuffix(stmt, ",")
}

func AllOrders(c *gin.Context) { 
	var ( 
		order  m.OrderModel
		orders []m.OrderModel
		detail  m.DetailPesanan
		details []m.DetailPesanan
		total  int
	)
	rows, err := conn.DB.Query("select pesanan.id_pesanan, pesanan.tglorder, pesanan.id_user, pesanan.total, pesanan.status, pesanan.image_order, pesanan.tglupload, user.name from pesanan INNER JOIN user ON pesanan.id_user = user.id_user;")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&order.IdPesanan, &order.Tglorder, &order.IdUser, &order.Total, &order.Status, &order.ImageOrder, &order.TglUpload, &order.Name )
		if err != nil {
			log.Fatal(err.Error())
		}
		
		
		details = nil
		if order.IdPesanan != 0 {
		
			rowsPesanan, err := conn.DB.Query(" SELECT detail_pesanan.*, product.price, product.name FROM detail_pesanan INNER JOIN product ON detail_pesanan.id_produk = product.id_product  where id_pesanan=" + strconv.Itoa(order.IdPesanan)+ ";")

			// SELECT `detail_pesanan`.*, `product`.`price` FROM `detail_pesanan` INNER JOIN `product` ON `detail_pesanan`.`id_produk` = `product`.`id_product` WHERE id_pesanan = '12'

			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(order.IdPesanan)

			for rowsPesanan.Next() {
				
				err = rowsPesanan.Scan(&detail.Id, &detail.IdProduk, &detail.IdPesanan, &detail.Quantity, &detail.Price, &detail.Name)
				if err != nil {
					log.Fatal(err.Error())
				}
				// fmt.Println(err)
				
				total =  detail.Price * detail.Quantity
				fmt.Println(detail)
				details = append(details, detail)
			}	

				order.DetailPesanan = details
				order.Total = total

				// order.Tglorder2 = order.Tglorder.Format("02-January-2006 15:04:05 WIB")
				
				
		}	
		// details = details[0:0]
		
		orders = append(orders, order) 

	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": orders,
		"count":  len(orders),
	})

}

func SelectOrder(c *gin.Context) {
	id := c.Param("data")
	var ( 
			order  m.OrderModel
			orders []m.OrderModel
			detail  m.DetailPesanan
			details []m.DetailPesanan
			total  int
	)
	
	// Cara prepare
	row, err := conn.DB.Query("select id_pesanan, tglorder, id_user, total, status, image_order, tglupload from pesanan where id_user = ?;", id)

	if err != nil {
		log.Fatal(err.Error())
	}
	for row.Next() {
		err := row.Scan(&order.IdPesanan, &order.Tglorder, &order.IdUser, &order.Total, &order.Status, &order.ImageOrder, &order.TglUpload)
		
		if err != nil {
			log.Fatal(err.Error())
		}
		details = nil
		if order.IdPesanan != 0 {
		
			rowsPesanan, err := conn.DB.Query(" SELECT detail_pesanan.*, product.price, product.name FROM detail_pesanan INNER JOIN product ON detail_pesanan.id_produk = product.id_product  where id_pesanan=" + strconv.Itoa(order.IdPesanan)+ ";")

			// SELECT `detail_pesanan`.*, `product`.`price` FROM `detail_pesanan` INNER JOIN `product` ON `detail_pesanan`.`id_produk` = `product`.`id_product` WHERE id_pesanan = '12'

			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(order.IdPesanan)

			for rowsPesanan.Next() {
				
				err = rowsPesanan.Scan(&detail.Id, &detail.IdProduk, &detail.IdPesanan, &detail.Quantity, &detail.Price, &detail.Name)
				if err != nil {
					log.Fatal(err.Error())
				}
				// fmt.Println(err)
				
				total =  detail.Price * detail.Quantity
				fmt.Println(detail)
				details = append(details, detail)
			}	

				order.DetailPesanan = details
				order.Total = total
				
		}	

		orders = append(orders, order) 
		
	}
	defer row.Close()

	// fmt.Println(order.Imageupload)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully get", id),
		"result": orders,
	})

}

func Produksell(c *gin.Context){
	var ( 
		order  m.SellsProduct
		orders []m.SellsProduct
	)
	rows, err := conn.DB.Query("SELECT id_produk, SUM(quantity) as jumlah, name,  ( SUM(quantity) * price) as total, image, price FROM detail_pesanan INNER JOIN product ON detail_pesanan.id_produk = product.id_product GROUP BY id_produk")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Jumlah, &order.Name, &order.Total, &order.Image, &order.Price )
		if err != nil {
			log.Fatal(err.Error())
		}

		orders = append(orders, order) 

	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": orders,
		"count":  len(orders),
	})
}


func Orderuploadimage(c *gin.Context) {
	// var path string
	form, _ := c.MultipartForm()
	files := form.File["order"]
	for _, file := range files {
		
		path := "./imageorder/"+file.Filename
		err := c.SaveUploadedFile(file,path)
		if err != nil {
			log.Fatal(err)
		}
		url := "http://localhost:5000/imageorder/" + file.Filename
		 
		fmt.Println(url)
		c.JSON(http.StatusOK, gin.H{
			"url": url,
			"name": file.Filename,
		})
	}
}
func UpdateImageOrder(c *gin.Context) {
	var buffer bytes.Buffer
	var Order m.OrderModel
	// var layoutFormat, value string
	// var date time.Time
	no := c.Param("no")
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Order, binding.JSON)
	fmt.Println(err)
	var now = time.Now()
	
	// layoutFormat = "2006-01-02 15:04:05"
	// value = date
	// date, _ = time.Parse(layoutFormat, value)
	
	stmt, err := conn.DB.Query("update pesanan set status=?, image_order=?, tglupload=? where id_pesanan =?",Order.Status, Order.ImageOrder, now, no)
	fmt.Println(now)
	if err != nil {
		log.Fatal(err.Error())
	}

	buffer.WriteString(Order.Status)
	buffer.WriteString(" ")
	buffer.WriteString(Order.ImageOrder)
	defer stmt.Close()
	status := buffer.String()
	fmt.Println(status)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", status),
		"tglupload":Order.TglUpload,
	})
}
func UbahDataOrder(c *gin.Context) {
	var buffer bytes.Buffer
	var Order m.OrderModel
	id := c.Param("id")
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Order, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Query("update pesanan set status=?, tglorder=?, total=? where id_pesanan =?",Order.Status, Order.Tglorder, Order.Total, id)

	if err != nil {
		log.Fatal(err.Error())
	}

	buffer.WriteString(Order.Status)
	buffer.WriteString(" ")
	// buffer.WriteString(Order.Tglorder)
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(Order.Total))
	defer stmt.Close()
	status := buffer.String()
	fmt.Println(status)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", status),
	})
}
func DeleteDataOrder(c *gin.Context) {
	id := c.Param("id")

	// Cara prepare
	stmt, err := conn.DB.Prepare("delete pesanan.*, detail_pesanan.* from pesanan, detail_pesanan where pesanan.id_pesanan = detail_pesanan.id_pesanan and pesanan.id_pesanan = ?")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully deleted", id),
	})

}
func ExportDataProdukSell(c *gin.Context) {
	tgl1 := c.Param("tgl")
	tgl2 := c.Param("tglb")
	var ( 
		Sell  m.ExSellsProduct
		Sells []m.ExSellsProduct
		layoutFormat,value,value2,tglfix,tglfix1 string
		date1,date,date2 time.Time
	)
	err := c.ShouldBindBodyWith(&Sell, binding.JSON)
	fmt.Println(err)
	rows, err := conn.DB.Query("SELECT pesanan.id_pesanan, pesanan.tglorder, product.name, detail_pesanan.quantity, product.price, (detail_pesanan.quantity*product.price) AS total_price FROM pesanan JOIN detail_pesanan ON pesanan.id_pesanan=detail_pesanan.id_pesanan JOIN product ON product.id_product=detail_pesanan.id_produk WHERE pesanan.tglorder BETWEEN '"+tgl1+"' AND '"+tgl2+"';")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&Sell.IdPesanan, &Sell.Tglorder, &Sell.Name, &Sell.Quantity, &Sell.Price, &Sell.TotalPrice)
		if err != nil {
			log.Fatal(err.Error())
		}
		
		Sells = append(Sells, Sell)
	}
	defer rows.Close()
	
	layoutFormat = "2006-01-02"
	value = tgl1
	date1, _ = time.Parse(layoutFormat, value)
	// tglfix = date1.String()
	tglfix = date1.Format("02, January 2006")

	value2 = tgl2
	date2, _ = time.Parse(layoutFormat, value2)
	tglfix1 = date2.Format("02, January 2006")


	f, err := excelize.OpenFile("./report/laporan.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	index := f.NewSheet("Sheet1")
	fmt.Println(Sells)
	var row = 8
	var angka = 0
	_ = f.InsertRow("Sheet1", row+1)
	f.SetCellValue("Sheet1", "B4", tglfix)
	f.SetCellValue("Sheet1", "B5", tglfix1)
	

	for _, v := range Sells {
		row++
		angka++
        f.SetCellValue("Sheet1", "A"+strconv.Itoa(row), strconv.Itoa(angka))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(row), v.IdPesanan)
		// layoutFormat = "02-01-2006"
		value = v.Tglorder
		date, _ = time.Parse(time.RFC3339, value)
		var dateS1 = date.Format("Monday 02, January 2006 15:04 WIB")
		
        f.SetCellValue("Sheet1", "C"+strconv.Itoa(row), dateS1)
        f.SetCellValue("Sheet1", "D"+strconv.Itoa(row), v.Name)
        f.SetCellValue("Sheet1", "E"+strconv.Itoa(row), v.Quantity)
        f.SetCellValue("Sheet1", "F"+strconv.Itoa(row), v.Price)
        f.SetCellValue("Sheet1", "G"+strconv.Itoa(row), v.TotalPrice)
    }
	f.SetActiveSheet(index)
	url := "/report/Laporan Data Produk Terjual " + tgl1 + " sampai " + tgl2 + ".xlsx"
	err = f.SaveAs("." + url)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": Sells,
		"count":  len(Sells),
		"url":  url,
	})
}