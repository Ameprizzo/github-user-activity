package utils

import (
	"fmt"
	"github-activity/internal"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Define styles using lipgloss
var (
	titleStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6"))
	eventTypeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")).Bold(true)
	repoStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Bold(true)
	refStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")).Italic(true)
	commitCountStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB86C")).Bold(true)
)

// DisplayEvents formats and displays the GitHub events in the terminal
func DisplayEvents(events []internal.Event) {
	// Maps to group events by type and repo
	pushEvents := make(map[string]int)
	createEvents := make(map[string]map[string][]string)
	deleteEvents := make(map[string]map[string][]string)
	genericEvents := make(map[string]map[string]int) // Group other events like WatchEvent, PullRequestEvent, etc.

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
			if genericEvents[event.Type] == nil {
				genericEvents[event.Type] = make(map[string]int)
			}
			genericEvents[event.Type][event.Repo.Name]++
		}
	}

	// Display the header
	fmt.Println(titleStyle.Render(fmt.Sprintf("Fetching activity for user: %s", events[0].Actor.Login)))

	// Display combined PushEvents
	for repo, count := range pushEvents {
		fmt.Println(eventTypeStyle.Render(fmt.Sprintf("- Pushed %d commit(s)", count)), "to", repoStyle.Render(repo))
	}

	// Display combined CreateEvents with specific names
	for repo, refTypes := range createEvents {
		for refType, refs := range refTypes {
			if refType == "repository" {
				fmt.Println(eventTypeStyle.Render("- Created a new repository:"), repoStyle.Render(repo))
			} else {
				fmt.Println(eventTypeStyle.Render(fmt.Sprintf("- Created %d new %s(s)", len(refs), refType)),
					refStyle.Render(fmt.Sprintf("(%s)", strings.Join(refs, ", "))), "in", repoStyle.Render(repo))
			}
		}
	}

	// Display combined DeleteEvents with specific names
	for repo, refTypes := range deleteEvents {
		for refType, refs := range refTypes {
			fmt.Println(eventTypeStyle.Render(fmt.Sprintf("- Deleted %d %s(s)", len(refs), refType)),
				refStyle.Render(fmt.Sprintf("(%s)", strings.Join(refs, ", "))), "in", repoStyle.Render(repo))
		}
	}

	// Display other events like WatchEvent, PullRequestEvent, etc.
	for eventType, repos := range genericEvents {
		for repo, count := range repos {
			fmt.Println(eventTypeStyle.Render(fmt.Sprintf("- %s", eventType)), repoStyle.Render(repo), commitCountStyle.Render(fmt.Sprintf("(%d times)", count)))
		}
	}
}
