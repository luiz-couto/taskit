package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"

	"github.com/spf13/cobra"
)

// URL to connect to webserver
var URL string = "http://localhost:49160"

// Task Struct
type Task struct {
	Rowid          int    `json:"rowid"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Priority       int    `json:"priority"`
	Blocked        int    `json:"blocked"`
	Deadline       string `json:"deadline"`
	WorkingEnter   string `json:"workingEnter"`
	WorkingElapsed string `json:"workingElapsed"`
	CreatedAt      string `json:"createdAt"`
	TimeEstimate   string `json:"timeEstimate"`
}

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "List all the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pFlag, _ := cmd.Flags().GetInt("priority")
		dFlag, _ := cmd.Flags().GetString("deadline")
		cFlag, _ := cmd.Flags().GetString("createdAt")
		tasksArr := getAllTasks(pFlag, dFlag, cFlag)
		drawBoard(tasksArr)
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)
	boardCmd.Flags().IntP("priority", "p", -1, "Filter tasks by priority")
	boardCmd.Flags().StringP("deadline", "d", "", "Filter tasks by deadline")
	boardCmd.Flags().StringP("createdAt", "c", "", "Filter tasks by created day (added time)")
}

// Get all tasks from localhost webserver
func getAllTasks(pFlag int, dFlag string, cFlag string) []Task {
	resp, err := http.Get(URL + "/tasks")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var tasksArray []Task
	e := json.NewDecoder(resp.Body).Decode(&tasksArray)
	if e != nil {
		fmt.Println(e)
	}

	if pFlag >= 0 {
		var filteredArr []Task
		for _, v := range tasksArray {
			if v.Priority == pFlag {
				filteredArr = append(filteredArr, v)
			}
		}
		return filteredArr
	}

	if dFlag != "" {
		if !verifyIfDateIsValid(dFlag) && dFlag != "today" {
			fmt.Println(`
			Date is not valid! Valid date is in format YYYY-MM-DD. 
			If you want to see tasks created today, you can also pass [today]
			`)
			os.Exit(0)
		} else {

			if dFlag == "today" {
				now := time.Now()
				nowAsString := now.Format("2006-01-02T15:04:05-0700")
				dFlag = nowAsString[0:10]
			}

			var filteredArr []Task
			for _, v := range tasksArray {
				if v.Deadline == dFlag {
					filteredArr = append(filteredArr, v)
				}
			}
			return filteredArr
		}
	}

	if cFlag != "" {

		if !verifyIfDateIsValid(cFlag) && cFlag != "today" {
			fmt.Println(`
			Date is not valid! Valid date is in format YYYY-MM-DD. 
			If you want to see tasks created today, you can also pass [today]
			`)
			os.Exit(0)
		} else {

			if cFlag == "today" {
				now := time.Now()
				nowAsString := now.Format("2006-01-02T15:04:05-0700")
				cFlag = nowAsString[0:10]
			}

			var filteredArr []Task
			for _, v := range tasksArray {
				if v.CreatedAt[0:10] == cFlag {
					filteredArr = append(filteredArr, v)
				}
			}
			return filteredArr
		}
	}

	return tasksArray
}

// Draw the board with all the columns
func drawBoard(tasksArray []Task) {

	var toDoArr []Task
	var workingArr []Task
	var doneArr []Task
	var closeArr []Task

	for _, v := range tasksArray {
		switch v.Status {
		case "ToDo":
			toDoArr = append(toDoArr, v)
		case "Working":
			workingArr = append(workingArr, v)
		case "Done":
			doneArr = append(doneArr, v)
		case "Closed(Completed)", "Closed(Non-Completed)":
			closeArr = append(closeArr, v)
		}
	}

	newTitle := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignLeft).
			SetText(text)
	}
	menu := newPrimitive(getColumnString("TO DO", toDoArr))
	main := newPrimitive(getColumnString("WORKING", workingArr))
	sideBar := newPrimitive(getColumnString("DONE", doneArr))
	completed := newPrimitive(getColumnString("CLOSED", closeArr))

	grid := tview.NewGrid().
		SetRows(1, 0).
		SetColumns(0, 0, 0, 0).
		SetBorders(true).
		AddItem(newTitle("My Dashboard (Remember that the columns are scrollable)"), 0, 0, 1, 4, 0, 0, false)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(main, 1, 1, 1, 1, 0, 0, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 0, false).
		AddItem(completed, 1, 3, 1, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// Get the Text to be written in each column
func getColumnString(columnName string, tasksArray []Task) string {
	var columnString string = "/// " + columnName + " ///\n\n"
	for _, v := range tasksArray {
		columnString = columnString + "--------------------------\n"
		columnString = columnString + "ID: " + strconv.Itoa(v.Rowid) + "\n"
		columnString = columnString + "Title: " + v.Title + "\n"
		columnString = columnString + "Description: " + v.Description + "\n"
		columnString = columnString + "Priority: " + strconv.Itoa(v.Priority) + "\n"

		if v.Deadline != "" {
			columnString = columnString + "Deadline: " + v.Deadline + "\n"
		}

		if v.TimeEstimate != "" {
			columnString = columnString + "Time Estimate: " + v.TimeEstimate + " hrs\n"
		}

		if v.Blocked != -1 {
			if v.Blocked == 0 {
				columnString = columnString + "(( BLOCKED ))\n"
			} else {
				columnString = columnString + "(( BLOCKED BY TASK " + strconv.Itoa(v.Blocked) + " ))" + "\n"
			}
		}

		if columnName == "CLOSED" {
			columnString = columnString + "Status: " + v.Status + "\n"
		}
		columnString = columnString + "---------------------------\n"
	}
	return columnString
}
