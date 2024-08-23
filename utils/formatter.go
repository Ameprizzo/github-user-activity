package utils

import (
	"fmt"
	"github-activity/internal"
)

// DisplayEvents formats and displays the GitHub events in the terminal
func DisplayEvents(events []internal.Event) {
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commit(s) to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		case "PullRequestEvent":
			if event.Payload.Action == "opened" {
				fmt.Printf("- Opened a new pull request in %s\n", event.Repo.Name)
			}
		case "CreateEvent":
			if event.Payload.RefType == "repository" {
				fmt.Printf("- Created a new repository: %s\n", event.Repo.Name)
			} else if event.Payload.RefType == "branch" {
				fmt.Printf("- Created a new branch %s in %s\n", event.Payload.Ref, event.Repo.Name)
			} else if event.Payload.RefType == "tag" {
				fmt.Printf("- Created a new tag %s in %s\n", event.Payload.Ref, event.Repo.Name)
			}
		case "DeleteEvent":
			if event.Payload.RefType == "branch" {
				fmt.Printf("- Deleted the branch %s in %s\n", event.Payload.Ref, event.Repo.Name)
			} else if event.Payload.RefType == "tag" {
				fmt.Printf("- Deleted the tag %s in %s\n", event.Payload.Ref, event.Repo.Name)
			}
		// Add more cases as needed for other event types
		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
	}
}
