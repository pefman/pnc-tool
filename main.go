package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
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

	var (
		gettags string
	)

	// Parse flags
	flag.StringVar(&gettags, "gettags", "", "fetches tags from a youtube videoid")

	// Parse command-line arguments
	flag.Parse()

	// Check if no flags are provided
	if flag.NFlag() == 0 {
		fmt.Println("nothing to do.")
		flag.Usage()
		os.Exit(1)
	}

	//lers get some tags
	if gettags != "" {
		client := &http.Client{
			Transport: &transport.APIKey{Key: config.APIKey},
		}

		service, err := youtube.New(client)
		if err != nil {
			log.Fatalf("Error creating YouTube client: %v", err)
		}

		call := service.Videos.List([]string{"snippet"}).Id(gettags)

		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error making videos API call: %v", err)
		}

		if len(response.Items) > 0 {
			video := response.Items[0]

			// Convert the tags to JSON
			tagsJSON, err := json.Marshal(video.Snippet.Tags)
			if err != nil {
				log.Fatalf("Error marshalling tags to JSON: %v", err)
			}

			fmt.Printf(string(tagsJSON))
		}
	}

	fmt.Println("")
}
