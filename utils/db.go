package utils

import (
	"context"
	"go-template/ent"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DbConn *ent.Client

func CreateDbConnection() {
	dsn := "host=localhost user=dya_only dbname=insta sslmode=disable"
	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to connect database. \n", err)
		os.Exit(1)
	}

	log.Println("DB Connected.")
	DbConn = client

	ctx := context.Background()
	if err := DbConn.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
