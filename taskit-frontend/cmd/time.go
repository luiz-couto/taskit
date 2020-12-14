package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "See the working time spent on a task",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Need to specify taskID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		task := getTaskByID(args[0])
		now := time.Now()

		parse, err := time.Parse("2006-01-02T15:04:05-0700", task.WorkingEnter)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		diff := now.Sub(parse)
		out := time.Time{}.Add(diff)
		fmt.Println(out.Format("15h 04m 05s"))

	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
