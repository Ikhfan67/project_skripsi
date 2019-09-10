package controller

import (
	"fmt"
	"log"
	"net/http"
	"bytes"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	"os"


	conn "skripsi_api/config"
	m "skripsi_api/models/products"
	n "skripsi_api/models/order"
	
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" 
)

func init() {
	conn.Connect()
}

func AllProducts(c *gin.Context) { 
	var ( 
		product  m.TransformedProduct
		products []m.TransformedProduct
	)
	rows, err := conn.DB.Query("select id_product, name, description, price,image from product;")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image)
		products = append(products, product) 
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// fmt.Println(product.Image)
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": products,
		"count":  len(products),
	})

}

func CreateProduct(c *gin.Context) {
	var buffer bytes.Buffer
	var Product m.ProductModel
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Product, binding.JSON)
	fmt.Println(err)
	stmt, err := conn.DB.Prepare("insert into product (name, description, price, image) values (?,?,?,?)")
	
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := stmt.Exec(Product.Name, Product.Description, Product.Price, Product.Image)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, _ := res.LastInsertId()
	buffer.WriteString(strconv.Itoa(int(id)))
	buffer.WriteString(" ")	
	buffer.WriteString(Product.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Product.Description)
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(Product.Price))
	buffer.WriteString(" ")
	buffer.WriteString(Product.Image)
	defer stmt.Close()
	produk := buffer.String()
	fmt.Println(produk)
	
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully add", produk),
		"id": id,
		"name": Product.Name,
		"description": Product.Description,
		"price": Product.Price,
		"image": Product.Image,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// Cara prepare
	stmt, err := conn.DB.Prepare("delete from product where id_product = ?")

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

func UpdateProduct(c *gin.Context) {
	var buffer bytes.Buffer
	var Product m.TransformedProduct
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Product, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update product set  name=?, description =?, price =?, image=? where id_product =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	Product.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Product.Name, Product.Description, Product.Price, Product.Image, Product.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	
	buffer.WriteString(Product.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Product.Description)
	buffer.WriteString(" ")
	buffer.WriteString(Product.Image)
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(Product.Price))
	defer stmt.Close()
	produk := buffer.String()
	fmt.Println(produk)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", produk),
	})
}

func UploadImage(c *gin.Context) {
	// var path string
	form, _ := c.MultipartForm()
	files := form.File["fileipan"]
	for _, file := range files {
		
		path := "./image/"+file.Filename
		err := c.SaveUploadedFile(file,path)
		if err != nil {
			log.Fatal(err)
		}
		url := "http://localhost:5000/image/" + file.Filename
		 
		fmt.Println(url)
		c.JSON(http.StatusOK, gin.H{
			"url": url,
			"name": file.Filename,
		})
	}
}
func Edituploadimage(c *gin.Context){
	form, _ := c.MultipartForm()
	files := form.File["editfileipan"]
	for _, file := range files {
		
		path := "./image/"+file.Filename
		err := c.SaveUploadedFile(file,path)
		if err != nil {
			log.Fatal(err)
		}
		url := "http://localhost:5000/image/" + file.Filename
		 
		fmt.Println(url)
		c.JSON(http.StatusOK, gin.H{
			"url": url,
			"name": file.Filename,
		})
	}
}
func DeleteImage(c *gin.Context) {
	path := c.Param("data")
	url := "./image/" + path
	fmt.Println(url)
	var err = os.Remove(url)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully deleted", url),
	})
}
func DeleteEditImage(c *gin.Context) {
	path := c.Param("data")
	url := "./image/" + path
	fmt.Println(url)
	var err = os.Remove(url)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully deleted", url),
	})
}
func DeleteUploadImage(c *gin.Context) {
	path := c.Param("data")
	url := "./imageorder/" + path
	fmt.Println(url)
	var err = os.Remove(url)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully deleted", url),
	})
}

func Test(c *gin.Context){
	fmt.Println("alo")
}

func DetailDataProdukSell(c *gin.Context) {
	var ( 
		detail  n.ExSellsProduct
		details  []n.ExSellsProduct
	)
	path := c.Param("name")
	rows, err := conn.DB.Query("SELECT pesanan.id_pesanan, pesanan.tglorder, product.name, detail_pesanan.quantity, product.price, (detail_pesanan.quantity*product.price) AS total_price FROM pesanan JOIN detail_pesanan ON pesanan.id_pesanan=detail_pesanan.id_pesanan JOIN product ON product.id_product=detail_pesanan.id_produk WHERE product.name = '"+path+"'")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
	err = rows.Scan(&detail.IdPesanan, &detail.Tglorder, &detail.Name, &detail.Quantity, &detail.Price, &detail.TotalPrice)
		if err != nil {
			log.Fatal(err.Error())
		}
		details = append(details, detail)
	}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"result": details,
		"count":  len(details),
	})
}


