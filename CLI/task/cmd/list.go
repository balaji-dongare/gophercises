package cmd

import (
	"fmt"

	"github.com/balaji-dongare/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

var listallTask = dbrepository.ReadTodosTaskFromDB

//ListTask  Show todo task list
var ListTask = &cobra.Command{
	Use:   "list",
	Short: "list is a CLI command to show todo task list ",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := listallTask()
		if err != nil {
			fmt.Print("\nNo task found")
		}
		if len(tasks) > 0 {
			for i, task := range tasks {
				fmt.Printf("%d. %v ", i+1, task)
			}
		} else {
			fmt.Print("No Tasks are in todos...Great!")
		}
	},
}

func init() {
	RootCmd.AddCommand(ListTask)
}
