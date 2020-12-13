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

	fmt.Println("Task updated successfully!")

}
