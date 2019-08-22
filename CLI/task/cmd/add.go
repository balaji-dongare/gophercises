package cmd

import (
	"fmt"
	"strings"

	"github.com/TestGit/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

//AddTask Add new task in todo task list
var AddTask = &cobra.Command{
	Use:   "add",
	Short: "add is a CLI command to add your todo  into task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		status, err := dbrepository.InsertTaskIntoDB(task)
		if err != nil {
			fmt.Println("Not able to add task to todo")
		}
		if status {
			fmt.Printf("Task:\"%s\" is Added in todo list\n", task)
		} else {
			fmt.Printf("Unable to add Task:\"%s\" in todo list\n", task)
		}
	},
}

func init() {
	RootCmd.AddCommand(AddTask)
}
