package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// Task Struct
type Task struct {
	Rowid       int    `json:"rowid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "List all the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)
}

// Get all tasks from localhost webserver
func getAllTasks() {
	resp, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var tasksArray []Task
	e := json.NewDecoder(resp.Body).Decode(&tasksArray)
	if e != nil {
		fmt.Println(e)
	}

	for _, v := range tasksArray {
		fmt.Printf("--------------------------\n")
		fmt.Printf("ID: %v\n", v.Rowid)
		fmt.Printf("Title: %v\n", v.Title)
		fmt.Printf("Description: %v\n", v.Description)
		fmt.Printf("--------------------------\n")
	}

}
