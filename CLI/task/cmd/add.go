package cmd

import (
	"fmt"
	"strings"

	"github.com/balaji-dongare/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

//AddTask Add new task in todo task list
var AddTask = &cobra.Command{
	Use:   "add",
	Short: "add is a CLI command to add your todo  into task list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			task := strings.Join(args, " ")
			status, err := dbrepository.InsertTaskIntoDB(task)
			if err != nil {
				fmt.Printf("Unable to add Task:\"%s\" in todo list\n", task)
			}
			if status {
				fmt.Printf("Task:\"%s\" is Added in todo list\n", task)
			}
		} else {
			fmt.Printf("Please provide task\n")
		}	
	},
}

func init() {
	RootCmd.AddCommand(AddTask)
}
