package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Config() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("Error pinging database connection: %v\n", err)
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	return db

}
