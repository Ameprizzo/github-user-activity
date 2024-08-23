package internal

import "time"

// Event represents a GitHub event
type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

// Actor represents the user who triggered the event
type Actor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// Repo represents the repository associated with the event
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload represents the payload of the event
type Payload struct {
	PushID      int      `json:"push_id"`
	Size        int      `json:"size"`
	Commits     []Commit `json:"commits"`
	Action      string   `json:"action"`
	Ref         string   `json:"ref"`         // Branch or tag name
	RefType     string   `json:"ref_type"`    // Type of ref (branch, tag)
	PusherType  string   `json:"pusher_type"` // Type of pusher
	PullRequest struct {
		URL string `json:"url"`
	} `json:"pull_request"`
}

// Commit represents a single commit in a PushEvent
type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
	URL     string `json:"url"`
}
