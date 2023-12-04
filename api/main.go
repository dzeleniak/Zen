package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load()
}

func main() {

	PORT := os.Getenv("PORT")

	USER := os.Getenv("DB_USER")
	PW := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, HOST, DBNAME)
	// username:password@protocol(address)/dbname?param=value

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Ping()

	e := echo.New()

	serverAddress := fmt.Sprintf(":%s", PORT)
	e.Logger.Fatal(e.Start(serverAddress))
}