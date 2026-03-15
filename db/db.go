// package db

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func InitDB() {

// 	var err error

// 	DB, err = sql.Open(
// 		"mysql",
// 		"root:123456@tcp(127.0.0.1:3306)/quanlylophoc",
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

dsn := os.Getenv("MYSQL")

	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(" Cannot connect database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(" Cannot ping database:", err)
	}

	log.Println(" Connected to Railway MySQL")
}