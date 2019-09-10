package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	co "skripsi_api/controller/auth"
	c "skripsi_api/controller/products"
	d "skripsi_api/controller/profile"
	e "skripsi_api/controller/komen"
	f "skripsi_api/controller/order"
	"github.com/gin-gonic/contrib/static" 
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./coffee-shop1/dist", true)))

	//Masalah 404
	r.NoRoute(func(c *gin.Context){
		c.File("./coffee-shop1/dist/index.html")
	})

	r.Use(static.Serve("/image", static.LocalFile("./image", true)))
	r.Use(static.Serve("/imageorder", static.LocalFile("./imageorder", true)))
	r.Use(static.Serve("/gambar_bank", static.LocalFile("./gambar_bank", true)))
	r.Use(static.Serve("/gambar_profile", static.LocalFile("./gambar_profile", true)))
	r.Use(static.Serve("/report", static.LocalFile("./report", true)))
	r.Use(cors.AllowAll())
	fmt.Println("Hayy")
	auth := r.Group("/auth")
	{
		auth.POST("/login", co.LoginHandler)
		auth.POST("/register", co.Register)
	}
	cek := r.Group("/login")
	{
		cek.GET("/cek", co.CekAdmin)
	}
	image:= r.Group("/image")
	{
		image.Use(co.Auth)
		{
			image.DELETE("/removeImage/:data", c.DeleteImage)
			image.DELETE("/removeEditImage/:data", c.DeleteEditImage)
		}
	}
	r.DELETE("/removeUploadImage/:data", c.DeleteUploadImage)
	r.POST("/profileuploadimage", d.Profileuploadimage)
	r.DELETE("/removeProfilImage/:data", d.DeleteProfilImage)
	r.PUT("/uploadprofilefix/:id", d.EditProfilImage)
	r.GET("/getimageprofile/:id", d.GetImageProfile)
	r.GET("/getdataprofile/:id", d.GetDataProfile)
	r.PUT("/updatedataprofile", d.EditDataProfil)
	r.POST("/tambahdataprofile", d.TambahDataProfil)
	r.POST("/test/:tgl"+"/:tglb", f.ExportDataProdukSell)
	product := r.Group("/product")
    {
		product.GET("/all/home", c.AllProducts)
		product.Use(co.Auth)
		{
		product.GET("/all", c.AllProducts)
		product.POST("/add", c.CreateProduct)
		product.POST("/uploadimage", c.UploadImage)
		product.POST("/edituploadimage", c.Edituploadimage)
		product.PUT("/edit/:id", c.UpdateProduct)
		product.DELETE("/:id", c.DeleteProduct)
		}
	}
	profile := r.Group("/profile")
	{
		profile.Use(co.Auth)
		{
			profile.GET("all", d.AllProfiles)
			profile.PUT("/edit/:id", d.UpdateProfile)
			profile.PUT("/reset/:id", d.ResetEmailPassword)
			profile.PUT("/new/:id", d.NewEmailPassword)
			profile.DELETE("/:id", d.DeleteProfile)
		}
	}
	komen := r.Group("/komentar")
	{
		komen.POST("/add", e.CreateKomen)
		komen.GET("/get", e.GetKomen)
		komen.DELETE("/delete/:id", e.DeleteKomen)
	}
	order := r.Group("/order")
	{
		order.POST("/add", f.CreateOrder)
		order.GET("/all", f.AllOrders)
		order.GET("/produksell", f.Produksell)
		order.GET("/get/:data", f.SelectOrder)
		order.POST("/orderuploadimage", f.Orderuploadimage)
		order.PUT("/edit/:no", f.UpdateImageOrder)
	}
	ubah := r.Group("/edit/order")
	{
		ubah.Use(co.Auth)
		{
			ubah.PUT("/:id", f.UbahDataOrder)
			ubah.DELETE("/delete/:id", f.DeleteDataOrder)
		}
	}
	detail := r.Group("/details")
	{
		detail.POST("/product/:name", c.DetailDataProdukSell)	
	}
	r.Run(":5000")
}
