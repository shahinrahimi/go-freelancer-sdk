package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/go-freelancer-sdk/v1"
)

func main() {

	// Create a new client
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	token := os.Getenv("FREELANCER_ACCESS_TOKEN")
	if token == "" {
		panic("FREELANCER_ACCESS_TOKEN is not set")
	}
	c := freelancer.NewClient(token)
	c.Debug = true
	projectListService := c.NewProjectService()
	res, err := projectListService.Do(context.Background())
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("Response: %s", res)

}
