package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var recentCount int

var listRecentCmd = &cobra.Command{
	Use:   "list-recent",
	Short: "List the most recently used commands",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to list the most recent `recentCount` commands
		fmt.Printf("Listing the most recent %d commands\n", recentCount)
	},
}

func init() {
	listRecentCmd.Flags().IntVarP(&recentCount, "number", "n", 10, "Number of commands to list")
	rootCmd.AddCommand(listRecentCmd)
}
