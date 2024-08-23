package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "github-activity",
    Short: "Fetch and display GitHub user activity",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide a GitHub username")
            os.Exit(1)
        }
        username := args[0]
        // Call your function to fetch and display activity here
        fmt.Println("Fetching activity for user:", username)
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
