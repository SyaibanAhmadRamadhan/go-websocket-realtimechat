package infra

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=root dbname=go-chat sslmode=disable")
	if err != nil {
		log.Fatalf("failed open connetion | err : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("failed ping to database | err : %v", err)
	}

	return db

}
