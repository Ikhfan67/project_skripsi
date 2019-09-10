package config
import (
	_ "github.com/go-sql-driver/mysql"
  	"database/sql"
	"fmt"
)

var DB *sql.DB
var err error

func Connect() (*sql.DB) {
  DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/skripsi_p?parseTime=true")
  checkErr(err)

	//defer db.Close()
	// make sure connection is available
	err = DB.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")
	return DB
}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}