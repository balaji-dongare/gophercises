package cmd

import (
	"fmt"
	"strconv"

	"github.com/balaji-dongare/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

var doTask = dbrepository.MarkTaskAsDone

//DoTask  Marks task as completed
var DoTask = &cobra.Command{
	Use:   "do",
	Short: "do is a CLI command to  mark task as completed ",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		if len(args) > 0 {
			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("Failed to parse Arg:", arg)
				} else {
					ids = append(ids, id)
				}
			}
			//fmt.Println(ids)
			notValidIds, taskdone, err := doTask(ids)
			note := ""
			if err != nil {
				note = "Sorry! Unable to mark as Complete."
				fmt.Printf("%v due to : %v", note, err)
			} else {
				if len(notValidIds) >= 1 {
					note = fmt.Sprintf("%v these ids not exist", notValidIds)
				} else {
					for i, task := range taskdone {
						fmt.Printf("%d. %v", i+1, task)
					}
				}
			}
		} else {
			fmt.Printf("Please provide task id")
		}
	},
}

func init() {
	RootCmd.AddCommand(DoTask)
}
