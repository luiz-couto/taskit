package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

		if task.Status == "Working" {
			now := time.Now()

			parse, err := time.Parse("2006-01-02T15:04:05-0700", task.WorkingEnter)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			diff := now.Sub(parse)

			elapsed, _ := strconv.ParseFloat(task.WorkingElapsed, 64)
			total := diff.Seconds() + elapsed

			totalAsString := fmt.Sprintf("%f", total)
			totalDuration, _ := time.ParseDuration(totalAsString + "s")

			out := time.Time{}.Add(totalDuration)
			fmt.Println(`
			Time working in this task:`, out.Format("15h 04m 05s"))
			fmt.Printf("\n")
		} else {

			elapsed, _ := strconv.ParseFloat(task.WorkingElapsed, 64)
			totalAsString := fmt.Sprintf("%f", elapsed)
			totalDuration, _ := time.ParseDuration(totalAsString + "s")
			out := time.Time{}.Add(totalDuration)
			fmt.Println(`
			Time working in this task:`, out.Format("15h 04m 05s"))
			fmt.Printf("\n")
		}

	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
