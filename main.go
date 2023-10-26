package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define the struct to match the JSON structure
type Author struct {
	Name string `json:"name"`
}

type Item struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	ContentText string   `json:"content_text"`
	Authors     []Author `json:"authors"`
}

type Feed struct {
	Version string `json:"version"`
	Title   string `json:"title"`
	HomeURL string `json:"home_page_url"`
	Items   []Item `json:"items"`
}

func main() {
	feedURL := "https://feeds.npr.org/1004/feed.json"

	resp, err := http.Get(feedURL)
	if err != nil {
		fmt.Printf("Error fetching the feed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the feed: %v\n", err)
		return
	}

	var feed Feed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		fmt.Printf("Error unmarshalling the JSON: %v\n", err)
		return
	}

	fmt.Println("Feed Title:", feed.Title)
	fmt.Println("Home URL:", feed.HomeURL)

	for _, item := range feed.Items {
		fmt.Println("\nItem Title:", item.Title)
		fmt.Println("Item URL:", item.URL)
		fmt.Println("Item Content:", item.ContentText)
		if len(item.Authors) > 0 {
			fmt.Println("Author:", item.Authors[0].Name)
		}
	}
}
