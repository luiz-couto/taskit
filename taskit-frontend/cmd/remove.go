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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Close a task",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Need to specify taskID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		removeTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeTask(taskID string) {

	task := getTaskByID(taskID)
	var val string

	if task.Status == "Done" {
		val = "Closed(Completed)"
	} else {
		val = "Closed(Non-Completed)"
	}

	requestBody, err := json.Marshal(map[string]string{
		"property": "status",
		"value":    val,
	})
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest(http.MethodPatch, URL+"/tasks/"+taskID, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Task closed!")
}

// Get task from localhost webserver
func getTaskByID(taskID string) Task {
	resp, err := http.Get(URL + "/tasks/" + taskID)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var tasksArray []Task
	e := json.NewDecoder(resp.Body).Decode(&tasksArray)
	if e != nil {
		fmt.Println(e)
	}

	if len(tasksArray) == 0 {
		fmt.Println("Task not found!")
		os.Exit(0)
	}

	return tasksArray[0]

}
