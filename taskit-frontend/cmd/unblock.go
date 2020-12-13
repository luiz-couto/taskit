package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock a task",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Need to specify taskID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskList := getAllTasks(-1, "")
		if !verifyIfTaskExists(args[0], taskList) {
			fmt.Println("Task not found!")
			os.Exit(0)
		}
		if !verifyIsTaskIsBlocked(args[0]) {
			fmt.Println("Task must be blocked to be unblocked!")
			os.Exit(0)
		}

		unblockTask(args[0])

	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}

func unblockTask(taskID string) {
	requestBody, err := json.Marshal(map[string]string{
		"property": "blocked",
		"value":    "-1",
	})
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/tasks/"+taskID, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Task unblocked successfully!")

}

func verifyIsTaskIsBlocked(taskID string) bool {
	task := getTaskByID(taskID)
	if task.Blocked == -1 {
		return false
	}
	return true
}
