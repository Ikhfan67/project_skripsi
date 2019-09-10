package auth
import(
    jwt "github.com/dgrijalva/jwt-go"

)
type Credential struct {
	Name  	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
	Role	 string `json:"role"`
}
type TransformedCredential struct {
	Id  	 int    `json:"id"`
	Name	 string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role	 string `json:"role"`
}
type MyClaims struct {
    Role    string   `json:"role"`
    jwt.StandardClaims
}