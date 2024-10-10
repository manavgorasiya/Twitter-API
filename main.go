package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

func main() {
	// Replace with your OAuth 1.0a credentials
	consumerKey := "BbxdHnwfTceNFOLMH9nt30jwA"
	consumerSecret := "Q5rnkRDqh99L6OIA0OfaHqWGS8zNhkybVKofBi8zkZ3HV2qsN9"
	accessToken := "1844366360580595723-PfoWFhviuk8GHvlOUPfZr2UHWldNi7"
	accessTokenSecret := "L47tQSh669r0f52cYsSXWVkGDjn2FajubVgdtdxwYS8pE"

	// OAuth1 authentication configuration
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create an HTTP client with the OAuth credentials
	httpClient := config.Client(oauth1.NoContext, token)

	// Post a new tweet
	tweetText := "Hello from Twitter API using OAuth 1.0a in Go!"
	postTweet(httpClient, tweetText)

	// Example: Deleting a tweet by its ID
	// tweetID := "1844468424316989824" // Replace with the actual Tweet ID to delete
	// deleteTweet(httpClient, tweetID)
}

// Function to post a new tweet using Twitter API v2
func postTweet(client *http.Client, tweetText string) {
	url := "https://api.twitter.com/2/tweets"
	tweet := map[string]string{"text": tweetText}

	jsonTweet, err := json.Marshal(tweet)
	if err != nil {
		log.Fatalf("Error marshalling tweet: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonTweet))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		// Read the response body for more information
		var respBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&respBody)
		log.Fatalf("Failed to post tweet: received status %s, response: %v", resp.Status, respBody)
	}

	fmt.Println("Successfully posted tweet!")
}

// Function to delete a tweet using Twitter API v2
// func deleteTweet(client *http.Client, tweetID string) {
// 	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

// 	req, err := http.NewRequest("DELETE", url, nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request: %v", err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Failed to delete tweet: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Handle both 200 OK and 204 No Content as successful responses
// 	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
// 		fmt.Println("Successfully deleted tweet!")
// 	} else {
// 		// Read the response body for more information
// 		log.Fatalf("Failed to delete tweet: received status %s", resp.Status)
// 	}
// }
