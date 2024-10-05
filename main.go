package freelancer

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	c := NewClient(token)
	projectListService := c.NewProjectService()
	res, err := projectListService.Do(context.Background())
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("Response: %s", res)

}
