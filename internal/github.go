package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.github.com/users/%s/events"

func FetchUserEvent(username string) ([]Event, error) {
	url := fmt.Sprintf(apiURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}
	fmt.Println(events)
	return events, nil
}
