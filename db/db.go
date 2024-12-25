package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func Connect() {
	var err error
	err = godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	urlDb := os.Getenv("DB_URL")

	DB, err = pgx.Connect(context.Background(), urlDb)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		fmt.Println("No se pudo conectar con la base de datos:")
	}

	fmt.Println("¡Conexión exitosa a PostgreSQL!")
}
