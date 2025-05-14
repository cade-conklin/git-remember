package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var popularCount int

var listPopularCmd = &cobra.Command{
	Use:   "list-popular",
	Short: "List the most popular commands",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to list the most popular `popularCount` commands
		fmt.Printf("Listing the most popular %d commands\n", popularCount)
	},
}

func init() {
	listPopularCmd.Flags().IntVarP(&popularCount, "number", "n", 10, "Number of commands to list")
	rootCmd.AddCommand(listPopularCmd)
}
