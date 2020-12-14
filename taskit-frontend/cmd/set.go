package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [taskID] [property] [new_value]",
	Short: "Update a task title/description",
	Long: `
	Set new value for a task property
		
	set [taskID] [property] [new_value]
		
	where property being one of:
		- title
		- description
		- status
		- priority
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("Need to specify taskID, property and new value")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		updateTaskValue(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func updateTaskValue(id, property, value string) {
	requestBody, err := json.Marshal(map[string]string{
		"property": property,
		"value":    value,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if property != "title" && property != "description" && property != "status" && property != "priority" && property != "deadline" {
		fmt.Println("property of task should be one of: title / description / status / priority / deadline")
		os.Exit(0)
	}

	if property == "status" && value != "ToDo" && value != "Working" && value != "Done" {
		fmt.Println("status should be one of: ToDo or Working or Done")
		os.Exit(0)
	}

	if property == "priority" {
		if _, err := strconv.Atoi(value); err != nil {
			fmt.Println("priority should be a integer!")
			os.Exit(0)
		} else if v, _ := strconv.Atoi(value); v < 0 {
			fmt.Println("priority should be positive or zero!")
			os.Exit(0)
		}
	}

	if property == "status" && value == "Done" {
		task := getTaskByID(id)
		if task.Blocked != -1 {
			fmt.Println("Blocked tasks cant be passed to Done!")
			os.Exit(0)
		}
	}

	if property == "deadline" {
		if value != "no-deadline" {
			if !verifyIfDateIsValid(value) {
				fmt.Println("Date is not valid! Valid date is in format YYYY-MM-DD")
				fmt.Println(`If you want to remove a deadline, just set it equal to "no-deadline"`)
				os.Exit(0)
			}
		} else {
			requestBody, _ = json.Marshal(map[string]string{
				"property": property,
				"value":    "",
			})
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/tasks/"+id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if property == "status" && value == "Done" {
		taskList := getAllTasks(-1, "")
		for _, t := range taskList {
			if strconv.Itoa(t.Blocked) == id {
				unblockTask(strconv.Itoa(t.Rowid))
				fmt.Println("Task " + strconv.Itoa(t.Rowid) + " is now unblocked!")
			}
		}
	}

	if property == "status" && value == "Working" {
		setWorkingEnter(id)
	}

	fmt.Println("Task updated successfully!")

}

func setWorkingEnter(id string) {

	now := time.Now()
	requestBody, err := json.Marshal(map[string]string{
		"property": "workingEnter",
		"value":    now.Format("2006-01-02T15:04:05-0700"),
	})

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/tasks/"+id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

}
