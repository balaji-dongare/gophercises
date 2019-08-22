package cmd

import (
	"fmt"
	"strconv"

	"github.com/TestGit/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

//DoTask  Marks task as completed
var DoTask = &cobra.Command{
	Use:   "do",
	Short: "do is a CLI command to  mark task as completed ",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse Arg:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		//fmt.Println(ids)
		notValidIds, taskdone, err := dbrepository.MarkTaskAsDone(ids)
		note := ""
		if err != nil {
			note = "Sorry! Unable to mark as Complete."
			fmt.Printf("\n%v due to : %v\n", note, err)
		} else {
			if len(notValidIds) >= 1 {
				note = fmt.Sprintf("\n%v these ids not exist\n", notValidIds)
			} else {
				fmt.Printf("Following Task Completed:\n")
				for i, task := range taskdone {
					fmt.Printf("%d. %v\n", i+1, task)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(DoTask)
}
