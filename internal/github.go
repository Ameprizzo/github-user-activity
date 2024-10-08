package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.github.com/users/%s/events"

// FetchUserActivity fetches the recent activity of a GitHub user
func FetchUserEvent(username string) ([]Event, error) {
	// Construct the API URL with the provided username
	url := fmt.Sprintf(apiURL, username)

	// Make an HTTP GET request to the GitHub API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("user not found. please check the username")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Parse the JSON response into a slice of Event structs
	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
