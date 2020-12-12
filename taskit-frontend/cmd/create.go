package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

		endTaskCreation(title, description)
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

func endTaskCreation(title string, description string) {
	requestBody, err := json.Marshal(map[string]string{
		"title":       title,
		"description": description,
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/tasks", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Task created successfully!")

}
