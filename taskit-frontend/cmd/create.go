package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a brand new task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var title string
		var description string
		var status string
		var priority string
		var deadline string
		var timeEstimate string

		for {
			if t, checkIfItsOk := readTitle(); checkIfItsOk {
				title = t
				break
			}
		}

		for {
			if d, checkIfItsOk := readDescription(); checkIfItsOk {
				description = d
				break
			}
		}

		for {
			if s, checkIfItsOk := readStatus(); checkIfItsOk {
				status = s
				break
			}
		}

		for {
			if p, checkIfItsOk := readPriority(); checkIfItsOk {
				priority = p
				break
			}
		}

		for {
			if dl, checkIfItsOk := readDeadline(); checkIfItsOk {
				deadline = dl
				break
			}
		}

		for {
			if te, checkIfItsOk := readTimeEstimate(); checkIfItsOk {
				timeEstimate = te
				break
			}
		}

		endTaskCreation(title, description, status, priority, deadline, timeEstimate)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func readTitle() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Task Name: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if len(text) == 0 {
		return "", false
	}
	return text, true
}

func readDescription() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Task Description: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text, true

}

func readStatus() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Initial Status (ToDo or Working): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text != "ToDo" && text != "Working" {
		return "", false
	}

	return text, true

}

func readPriority() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Task Priority (default to 0): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text, true
}

func readDeadline() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Task Deadline (YYYY-MM-DD - default to no-deadline): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if !verifyIfDateIsValid(text) {
		return "", false
	}

	return text, true
}

func readTimeEstimate() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> Time Estimate (In hours - Ex. 1, 1.5, 2, 2.2): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if !verifyIfFloatIsValid(text) {
		return "", false
	}

	return text, true
}

func verifyIfDateIsValid(date string) bool {
	if date == "" {
		return true
	}
	const layoutISO = "2006-01-02"
	_, err := time.Parse(layoutISO, date)
	if err != nil {
		return false
	}
	return true
}

func verifyIfFloatIsValid(num string) bool {
	if num == "" {
		return true
	}

	_, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return false
	}
	return true

}

func endTaskCreation(title string, description string, status string, priority string, deadline string, timeEstimate string) {

	var we string = ""

	if status == "Working" {
		now := time.Now()
		we = now.Format("2006-01-02T15:04:05-0700")
	}

	createdAt := time.Now()
	createdAtString := createdAt.Format("2006-01-02T15:04:05-0700")

	requestBody, err := json.Marshal(map[string]string{
		"title":        title,
		"description":  description,
		"status":       status,
		"priority":     priority,
		"deadline":     deadline,
		"workingEnter": we,
		"createdAt":    createdAtString,
		"timeEstimate": timeEstimate,
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(URL+"/tasks", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Task created successfully!")

}
