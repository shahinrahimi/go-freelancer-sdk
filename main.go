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
	c.SetBaseUrl(freelancer.BaseAPIMainURL)
	c.Debug = true

	pas := c.NewListActiveProjectsService()
	pas.SetQuery("python json")
	pas.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectTypeFixed})
	pas.SetProjectUpgrades([]freelancer.ProjectUpgradeType{freelancer.ProjectUpgradeTypeSealed})
	pas.SetFullDescription(true)
	pas.SetUserCountryDetails(true)
	pas.SetUserLocationDetails(true)
	res, err := pas.Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
	}
	for pr, p := range res.Result.Projects {
		log.Printf("Project #%d: %s", pr, p.Title)
		log.Printf("Project #%d: %s", pr, p.SeoURL)
		//log.Printf("Project #%d: %s", pr, p.Description)
		log.Printf("Project #%d: %s", pr, p.PreviewDesc)
		log.Printf("Project #%d: %s", pr, p.Location.City)
	}

	//log.Printf("Response: %s", res)

}
