package controller

import (
	"bytes"
	"time"
	"log"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	conn "skripsi_api/config"
    a "skripsi_api/models/auth"
    jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" 
) 

func init() {
	conn.Connect()
}

func LoginHandler(c *gin.Context) {
	var user a.TransformedCredential
	err := c.Bind(&user)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}

		// ini untuk ambil 1 query aja, karna user login dengan email dan pass yang sama cuman 1
		// tadi masalah karna penggunaan email ='?' =>  ini dianggapnya string bukan suatu parameter yang harus dimasukin ke golangnya 

		rows := conn.DB.QueryRow("select id_user, name, email, password, role from user where email= ?  and password= ? ", user.Email, user.Password)
		
		
			err = rows.Scan(&user.Id, &user.Name,&user.Email, &user.Password, &user.Role)
			
		
		// err = rows.Scan(&user.Email, &user.Password)
		if err != nil {
				// Kalau ini dipakai dia ngeberhentiin proses kalo gagal, karna ini hasilnya di kasih ke vue jadi harus returnnya berupa hasil juga
				// log.Fatal(err.Error())
				c.JSON(http.StatusUnauthorized, gin.H{
					"message":"Email atau Password tidak cocok",
				})		
		}else{
			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := sign.Claims.(jwt.MapClaims)
			claims["id_user"] = user.Id			
			claims["email"] = user.Email
			claims["password"] = user.Password
			claims["role"] = user.Role
			claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			token, err := sign.SignedString([]byte("secret"))
			email := user.Email
			name := user.Name
			id := user.Id
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			c.JSON(http.StatusOK, gin.H{
				"access_token": token,
				"access_email": email,
				"access_name": name,
				"access_id": id,
			})	
		}
}

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func Register(c *gin.Context) {
	var buffer bytes.Buffer
	var User a.TransformedCredential
	// pakai buffer agar pemrosesannya lebih cepat

	// name := c.PostForm("name")
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	err := c.ShouldBindBodyWith(&User, binding.JSON)
	// fmt.Println(fname)
	stmt, err := conn.DB.Prepare("insert into user (name, email, password) values (?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(User.Name, User.Email, User.Password)

	if err != nil {
		log.Fatal(err.Error())
	}

	buffer.WriteString(User.Name)
	buffer.WriteString(" ")
	buffer.WriteString(User.Email)
	buffer.WriteString(" ")
	buffer.WriteString(User.Password)

	defer stmt.Close()
	regist := buffer.String()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully registered", regist),
	})
}

func CekAdmin(c *gin.Context){
	tokenLogin := c.Request.Header.Get("Authorization")
	// claims := jwt.MapClaims{}
	claims := a.MyClaims{}
	token, err := jwt.ParseWithClaims(tokenLogin, &claims, func(token *jwt.Token) (interface{}, error) {
		// if jwt.GetSigningMethod("HS256") != token.Method {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }

		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println("gagal parsing ", err)
	}
	

	log.Println(token.Valid, claims)
	// var role a.MyClaims
	// roles:= role.Role

	if claims.Role != "admin" {
		result := gin.H{
			"message": "bukan admin",
			"error":   true,
		}
		c.JSON(http.StatusUnauthorized, result)
	}else {
		c.JSON(http.StatusOK, gin.H{
			"message": "admin",
			"error":   false,
		})
	}


}