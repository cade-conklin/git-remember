package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pageSize int

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all commands (paginated)",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to list all commands with pagination
		fmt.Printf("Listing all commands with page size %d\n", pageSize)
	},
}

func init() {
	listAllCmd.Flags().IntVarP(&pageSize, "page-size", "p", 20, "Number of commands per page")
	rootCmd.AddCommand(listAllCmd)
}
