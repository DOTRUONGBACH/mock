package database

import (
	"context"
	"fmt"
	"mock-project/ent"
)

func NewConnect() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=flight password=postgress sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return client, nil
}
