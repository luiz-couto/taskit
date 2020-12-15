package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block [taskID]",
	Short: "Block task",
	Long: `
	Using the command "block", you can block any task.
	You can simply block a task or you can also pass
	another task B that blocks this task. If B is 
	passed to Done, then the task is unblocked 
	automatically. You can also unblock any task at
	any time using the command "unblock"
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Need to specify taskID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskList := getAllTasks(-1, "", "")
		if !verifyIfTaskExists(args[0], taskList) {
			fmt.Println("Task not found!")
			os.Exit(0)
		}
		if !verifyIsTaskIsValid(args[0]) {
			fmt.Println(`
			Task must be with status ToDo or Working to be blocked!. 
			`)
			os.Exit(0)
		}

	OuterLoop:
		for {
			if v, checkIfItsOk := readTaskDependency(); checkIfItsOk {
				if v == 1 {

					for {
						otherTaskID := readOtherTaskID()
						if !verifyIfTaskExists(otherTaskID, taskList) {
							fmt.Println(`
							Task not found!
							`)
							break
						}

						if !verifyIsTaskIsValid(otherTaskID) {
							fmt.Println(`
							Task must be with status ToDo or Working to block other tasks!
							`)
							os.Exit(0)
						}

						value, _ := strconv.Atoi(args[0])
						taskTwoValue, _ := strconv.Atoi(otherTaskID)
						updateBlockedValue(value, taskTwoValue)
						break OuterLoop

					}

				} else {
					value, _ := strconv.Atoi(args[0])
					updateBlockedValue(value, 0)
					break
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}

func verifyIfTaskExists(taskID string, taskList []Task) bool {
	if _, err := strconv.Atoi(taskID); err != nil {
		fmt.Println("invalid TaskID!")
		os.Exit(0)
	}

	var found bool = false
	for _, v := range taskList {
		if i, _ := strconv.Atoi(taskID); i == v.Rowid {
			found = true
		}
	}
	return found
}

func verifyIsTaskIsValid(taskID string) bool {
	task := getTaskByID(taskID)
	if task.Status != "ToDo" && task.Status != "Working" {
		return false
	}
	return true
}

func readTaskDependency() (int, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Is blocked by other task? (y/n): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text != "y" && text != "n" {
		return -1, false
	}

	if text == "y" {
		return 1, true
	}

	return 2, true

}

func readOtherTaskID() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Task that is blocking this task: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func updateBlockedValue(taskID int, otherTaskID int) {

	requestBody, err := json.Marshal(map[string]string{
		"property": "blocked",
		"value":    strconv.Itoa(otherTaskID),
	})

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest(http.MethodPatch, URL+"/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Task blocked!")
}
