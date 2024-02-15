package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Config struct {
	APIKey    string `json:"APIKey"`
	ChannelId string `json:"ChannelId"`
}

func main() {
	// Read JSON file
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Parse JSON data into Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(config.APIKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// Call the Search.List method to retrieve the last 10 videos for a specific channel
	call := service.Search.List([]string{"snippet"}).
		ChannelId(config.ChannelId).
		MaxResults(10).
		Order("date")

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	// Print the video IDs of the last 10 videos
	for _, item := range response.Items {
		fmt.Println(item.Id.VideoId)
	}

}
