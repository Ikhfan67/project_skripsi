package controller
import (
	
	conn "skripsi_api/config"
	m "skripsi_api/models"
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin/binding"
	"log"
	"strconv"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func Init()  {
	conn.Connect()
}

func AllProfiles(c *gin.Context) { 
	var ( 
		profile  m.TransformedProfil
		profiles []m.TransformedProfil
	)
	rows, err := conn.DB.Query("select re.id_user, re.name, re.email, re.password, re.role, phone, address, postcode from profile AS us INNER JOIN user AS re ON re.id_user = us.id_user;")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&profile.Id, &profile.Name, &profile.Email, &profile.Password, &profile.Role, &profile.Phone, &profile.Address, &profile.Postcode)
		profiles = append(profiles, profile) 
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": profiles,
		"count":  len(profiles),
	})

}

func UpdateProfile(c *gin.Context) {
	var buffer bytes.Buffer
	var Profile m.TransformedProfil
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profile, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update profile AS us JOIN user AS re ON re.id_user = us.id_user SET re.name=?, re.role=?, us.phone =?, us.address =?, us.postcode =? where us.id_user =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	Profile.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Profile.Name, Profile.Role, Profile.Phone, Profile.Address, Profile.Postcode, Profile.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	
	buffer.WriteString(Profile.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Phone)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Address)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Postcode)
	defer stmt.Close()
	profile := buffer.String()
	fmt.Println(profile)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", profile),
	})
}


func ResetEmailPassword(c *gin.Context) {
	var buffer bytes.Buffer
	var Profile m.TransformedProfil
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profile, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update user SET email=?, password =? where id_user =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	Profile.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Profile.Email, Profile.Password, Profile.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	
	buffer.WriteString(Profile.Email)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Password)
	defer stmt.Close()
	profile := buffer.String()
	fmt.Println(profile)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", profile),
	})
}

func NewEmailPassword(c *gin.Context) {
	var buffer bytes.Buffer
	var Profile m.TransformedProfil
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profile, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update user SET name=?, email=?, password =? where id_user =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	Profile.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Profile.Name, Profile.Email, Profile.Password, Profile.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	buffer.WriteString(Profile.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Email)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Password)
	defer stmt.Close()
	profile := buffer.String()
	fmt.Println(profile)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", profile),
	})
}

func DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	// Cara prepare
	stmt, err := conn.DB.Prepare("delete profile.*, user.* from profile, user where profile.id_user = user.id_user and profile.id_user = ?")

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
func Profileuploadimage(c *gin.Context) {
	// var path string
	form, _ := c.MultipartForm()
	files := form.File["profile"]
	for _, file := range files {
		
		path := "./gambar_profile/"+file.Filename
		err := c.SaveUploadedFile(file,path)
		if err != nil {
			log.Fatal(err)
		}
		url := "http://localhost:5000/gambar_profile/" + file.Filename
		 
		fmt.Println(url)
		c.JSON(http.StatusOK, gin.H{
			"url": url,
			"name": file.Filename,
		})
	}
}
func DeleteProfilImage(c *gin.Context) {
	path := c.Param("data")
	url := "./gambar_profile/" + path
	fmt.Println(url)
	var err = os.Remove(url)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully deleted", url),
	})
}
func EditProfilImage(c *gin.Context){
	var buffer bytes.Buffer
	var Profile m.TransformedProfil
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profile, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update profile SET image=? where id_user =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	Profile.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Profile.Gambar,  Profile.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	buffer.WriteString(strconv.Itoa(Profile.Id))
	defer stmt.Close()
	profile := buffer.String()
	fmt.Println(profile)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully upload", profile),
	})

}
func GetImageProfile(c *gin.Context){
	path := c.Param("id")
	fmt.Println(path)
	
	var ( 
		profile  m.TransformedProfila
		profiles []m.TransformedProfila
	)
	rows, err := conn.DB.Query("select id_user, image from profile where id_user = "+path+";")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&profile.Id, &profile.Gambar)
		profiles = append(profiles, profile) 
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"result": profiles,
		"count":  len(profiles),
	})
}
func GetDataProfile(c *gin.Context){
	path := c.Param("id")
	fmt.Println(path)
	var ( 
		profile  m.TransformedProfila
		profiles []m.TransformedProfila
	)
	rows, err := conn.DB.Query("select re.name, re.email, re.password, phone, address, postcode from profile AS us INNER JOIN user AS re ON re.id_user = us.id_user where us.id_user=?",path)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&profile.Name, &profile.Email, &profile.Password, &profile.Phone, &profile.Address, &profile.Postcode)
		profiles = append(profiles, profile) 
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": profiles,
		"count":  len(profiles),
	})
}
func EditDataProfil(c *gin.Context){
	// path := c.Param("id")
	// fmt.Println(path)
	var buffer bytes.Buffer
	var Profile m.TransformedProfila
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profile, binding.JSON)
	fmt.Println(err)

	stmt, err := conn.DB.Prepare("update profile AS us JOIN user AS re ON re.id_user = us.id_user SET re.name=?, re.email=?, re.password=?, us.phone =?, us.address =?, us.postcode =? where us.id_user =?")

	if err != nil {
		log.Fatal(err.Error())
	}

	// Profile.Id, err = strconv.Atoi(c.Param("id"))

	_, err = stmt.Exec(Profile.Name, Profile.Email, Profile.Password, Profile.Phone, Profile.Address, Profile.Postcode, Profile.Id)

	if err != nil {
		log.Fatal(err.Error())
	}

	
	buffer.WriteString(Profile.Name)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Phone)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Address)
	buffer.WriteString(" ")
	buffer.WriteString(Profile.Postcode)
	defer stmt.Close()
	profile := buffer.String()
	fmt.Println(profile)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully edit", profile),
	})
}
func TambahDataProfil(c *gin.Context) {
	var buffer bytes.Buffer
	var Profil m.TransformedProfila
	// pakai buffer agar pemrosesannya lebih cepat
	
	// title := c.PostForm("title")
	// completed := c.PostForm("completed")
	err := c.ShouldBindBodyWith(&Profil, binding.JSON)
	fmt.Println(err)
	stmt, err := conn.DB.Prepare("insert into profile (id_user, phone, address, postcode) values (?,?,?,?)")
	
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := stmt.Exec(Profil.Id, Profil.Phone, Profil.Address, Profil.Postcode)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, _ := res.LastInsertId()
	buffer.WriteString(strconv.Itoa(int(id)))
	buffer.WriteString(" ")	
	buffer.WriteString(strconv.Itoa(Profil.Id))
	defer stmt.Close()
	produk := buffer.String()
	fmt.Println(produk)
	
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully add", produk),
		"id": id,
	})
}