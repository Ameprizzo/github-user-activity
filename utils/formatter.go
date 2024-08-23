// utils/formatter.go

package utils

import (
	"fmt"
	"github-activity/internal"
	"strings"
)

// DisplayEvents formats and displays the GitHub events in the terminal
func DisplayEvents(events []internal.Event) {
	// Maps to group events by type and repo
	pushEvents := make(map[string]int)
	createEvents := make(map[string]map[string][]string)
	deleteEvents := make(map[string]map[string][]string)

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			pushEvents[event.Repo.Name] += len(event.Payload.Commits)
		case "CreateEvent":
			if createEvents[event.Repo.Name] == nil {
				createEvents[event.Repo.Name] = make(map[string][]string)
			}
			createEvents[event.Repo.Name][event.Payload.RefType] = append(createEvents[event.Repo.Name][event.Payload.RefType], event.Payload.Ref)
		case "DeleteEvent":
			if deleteEvents[event.Repo.Name] == nil {
				deleteEvents[event.Repo.Name] = make(map[string][]string)
			}
			deleteEvents[event.Repo.Name][event.Payload.RefType] = append(deleteEvents[event.Repo.Name][event.Payload.RefType], event.Payload.Ref)
		default:
			// Other events are not combined and are displayed as is
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
	}

	// Display combined PushEvents
	for repo, count := range pushEvents {
		fmt.Printf("- Pushed %d commit(s) to %s\n", count, repo)
	}

	// Display combined CreateEvents with specific names
	for repo, refTypes := range createEvents {
		for refType, refs := range refTypes {
			if refType == "repository" {
				fmt.Printf("- Created a new repository: %s\n", repo)
			} else {
				fmt.Printf("- Created %d new %s(s) (%s) in %s\n", len(refs), refType, strings.Join(refs, ", "), repo)
			}
		}
	}

	// Display combined DeleteEvents with specific names
	for repo, refTypes := range deleteEvents {
		for refType, refs := range refTypes {
			fmt.Printf("- Deleted %d %s(s) (%s) in %s\n", len(refs), refType, strings.Join(refs, ", "), repo)
		}
	}
}
