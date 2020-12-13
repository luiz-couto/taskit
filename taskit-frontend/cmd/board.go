package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rivo/tview"

	"github.com/spf13/cobra"
)

// Task Struct
type Task struct {
	Rowid       int    `json:"rowid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "List all the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tasksArr := getAllTasks()
		drawBoard(tasksArr)
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)
}

// Get all tasks from localhost webserver
func getAllTasks() []Task {
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

	return tasksArray

}

// Draw the board with all the columns
func drawBoard(tasksArray []Task) {
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
	menu := newPrimitive(getColumnString("TO DO", tasksArray))
	main := newPrimitive(getColumnString("WORKING", tasksArray))
	sideBar := newPrimitive(getColumnString("COMPLETED", tasksArray))
	completed := newPrimitive(getColumnString("CLOSE", tasksArray))

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
		columnString = columnString + "Status: " + v.Status + "\n"
		columnString = columnString + "---------------------------\n"
	}
	return columnString
}
