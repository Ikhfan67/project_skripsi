package controller

import(
	conn "skripsi_api/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"bytes"
	"fmt"
	"log"
	"net/http"
	
	m "skripsi_api/models/komen"

)
func Init()  {
	conn.Connect()
}

func CreateKomen(c *gin.Context) {
	var buffer bytes.Buffer
	var Komen m.KomenModel
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Komen, binding.JSON)
	fmt.Println(err)
	stmt, err := conn.DB.Prepare("insert into komentar (name, email, phone, message) values (?,?,?,?)")
	
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(Komen.Name, Komen.Email, Komen.Phone, Komen.Message)

	if err != nil {
		log.Fatal(err.Error())
	}

	
	buffer.WriteString(Komen.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Komen.Email)
	buffer.WriteString(" ")
	buffer.WriteString(Komen.Phone)
	defer stmt.Close()
	komen := buffer.String()
	fmt.Println(komen)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully add", komen),
	})
}
func GetKomen(c *gin.Context) {
	var ( 
		Komen  m.TransformedKomenModel
		Komens []m.TransformedKomenModel
	)
	rows, err := conn.DB.Query("select id, name, email, phone,message from komentar;")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&Komen.Id, &Komen.Name, &Komen.Email, &Komen.Phone, &Komen.Message)
		Komens = append(Komens, Komen) 
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// fmt.Println(product.Image)
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": Komens,
		"count":  len(Komens),
	})
}
func DeleteKomen(c *gin.Context){
	id := c.Param("id")

	// Cara prepare
	stmt, err := conn.DB.Prepare("delete from komentar where id = ?")

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