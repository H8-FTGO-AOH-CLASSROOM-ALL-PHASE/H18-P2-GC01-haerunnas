package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Peringatan: Berkas .env tidak ditemukan")
	}

	database_url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("mysql", database_url)
	if err != nil {
		return nil, fmt.Errorf("Error Opening DB: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to Connect DB: %w", err)
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}
