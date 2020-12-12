package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "List all the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)
}

func listAllTasks() {
	fmt.Println("board called")
}
