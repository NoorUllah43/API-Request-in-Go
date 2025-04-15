package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgresqlDB() {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("database not connecting error", err)
	}

	fmt.Println("Database connected")
	
	defer db.Close()
}
