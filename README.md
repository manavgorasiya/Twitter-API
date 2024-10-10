# Twitter API Integration with Go

## Introduction
This project demonstrates how to interact with Twitter’s API using the Go programming language and OAuth 1.0a for user context authentication. The program covers two main functionalities:
1. Posting a new tweet.
2. Deleting an existing tweet.

The purpose of this assignment is to teach you how to authenticate with the Twitter API using OAuth 1.0a, make requests to the Twitter API endpoints, and handle API responses and errors effectively.

## Setup Instructions

### Setting up a Twitter Developer Account
To interact with the Twitter API, you'll need to create a Twitter Developer account and obtain the necessary credentials (API keys and tokens).

1. Go to the [Twitter Developer Platform](https://developer.twitter.com/en/portal/dashboard) and sign in with your Twitter account.
2. Apply for a Twitter Developer Account by following the instructions.
3. Once your application is approved, create a new project and app.
4. Navigate to **Projects & Apps → Your App → Keys and Tokens**.

### Generating API Keys and Access Tokens
1. In the **Keys and Tokens** section, generate the following credentials:
   - **API Key**
   - **API Secret Key**
   - **Access Token**
   - **Access Token Secret**
2. Ensure that your app’s **User Authentication Settings** have **OAuth 1.0a** enabled and permissions set to **Read and Write**.
3. You will use these credentials in your Go program to authenticate API requests.

### Running the Program
1. Clone or download this repository.
2. Install the necessary dependencies using Go modules:
   ```bash
   go mod init twitter-api-go
   go get github.com/dghubble/oauth1
   ```
3. Replace the placeholders for your Twitter API credentials in the `main.go` file:
   ```go
   consumerKey := "YOUR_API_KEY"
   consumerSecret := "YOUR_API_SECRET_KEY"
   accessToken := "YOUR_ACCESS_TOKEN"
   accessTokenSecret := "YOUR_ACCESS_TOKEN_SECRET"
   ```
4. To run the program:
   ```bash
   go run main.go
   ```

## Program Details

### Posting a New Tweet
The program uses the Twitter API v2's `/tweets` endpoint to post a tweet. The `postTweet` function sends a POST request with the tweet content in JSON format.

Example API Request:
```http
POST https://api.twitter.com/2/tweets
Authorization: OAuth 1.0a User Context
Content-Type: application/json

{
  "text": "Hello from Twitter API using Go!"
}
```

Example Response (Success):
```json
{
  "data": {
    "id": "1234567890123456789",
    "text": "Hello from Twitter API using Go!"
  }
}
```

### Deleting an Existing Tweet
The program deletes a tweet by sending a DELETE request to the `/2/tweets/:id` endpoint. The `deleteTweet` function takes the tweet’s ID as input and deletes it.

Example API Request:
```http
DELETE https://api.twitter.com/2/tweets/1234567890123456789
Authorization: OAuth 1.0a User Context
```

Example Response (Success):
- Status Code: `200 OK` or `204 No Content`.

### Example Commands
- **Post a new tweet**:
   Run the program and it will automatically post a tweet with the predefined text in the `main.go` file.
- **Delete a tweet**:
   Replace the `tweetID` in the `deleteTweet` function with the ID of the tweet you want to delete and run the program.

## Error Handling
The program includes error handling for common issues such as:
- **Invalid OAuth credentials**: The program checks for issues with OAuth tokens and provides appropriate error messages if authentication fails.
- **Forbidden actions**: If the app lacks proper permissions (e.g., not having write access), the program will return a `403 Forbidden` error with details.
- **Invalid tweet IDs**: If the tweet ID used for deletion is invalid, the program will notify the user with an appropriate error message.
- **Rate limiting**: If the API limit is reached, the program handles it by capturing the error and printing a meaningful message.
  
For example, if the program receives a `403 Forbidden` error when trying to post a tweet, the program logs:
```
Failed to post tweet: received status 403 Forbidden, response: map[detail:Your client app is not configured with the appropriate oauth1 app permissions for this endpoint.]
```

## Conclusion
This project demonstrates the basics of using the Twitter API to post and delete tweets with Go. It also highlights the importance of proper OAuth authentication and error handling when interacting with external APIs.