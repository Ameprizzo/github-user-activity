package cmd

import (
	"fmt"
	"github-activity/internal"
	"github-activity/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Fetch and display GitHub user activity",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a GitHub username")
			os.Exit(1)
		}
		username := args[0]

		fmt.Println("Fetching activity for user:", username)
		// Fetch the user's activity from GitHub
		events, err := internal.FetchUserEvent(username)
		if err != nil {
			fmt.Println("Error fetching user activity:", err)
			return
		}
		// Display the fetched events in the terminal
		utils.DisplayEvents(events)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
