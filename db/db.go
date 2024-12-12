package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgresql://user:user@localhost:5432/safehaven-db?sslmode=disable")

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
