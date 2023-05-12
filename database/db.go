package database

import (
	"context"
	"log"
	"mock-project/ent"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
)

func NewClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=flight password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
