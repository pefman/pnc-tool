package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func main() {

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Parse JSON data into Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Create a new YouTube service with the API key
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(config.APIKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	// ID of the video you want to fetch details for
	videoID := "stF2UxnLDOQ"

	// Make API call to get video details
	videoResponse, err := service.Videos.List([]string{"snippet"}).Id(videoID).Do()
	if err != nil {
		log.Fatalf("Error fetching video details: %v", err)
	}

	// Print the video details
	for _, item := range videoResponse.Items {
		fmt.Printf("Title: %s\n", item.Snippet.Title)
		fmt.Printf("Channel: %s\n", item.Snippet.ChannelTitle)
		fmt.Printf("Description: %s\n", item.Snippet.Description)
		fmt.Printf("Published At: %s\n", item.Snippet.PublishedAt)
	}
}
