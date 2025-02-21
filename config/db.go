package config

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connString := "server=localhost;database=testbtpns;user id=sa;password=password123!;encrypt=disable"
	
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err.Error())
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error testing database connection:", err.Error())
	}

	fmt.Println("Successfully connected to MSSQL database!")
}
