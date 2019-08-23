package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd  RootCommand
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI command to handle your todo task list",
}

// Execute is used to add command under root command
func execute() error {
	return RootCmd.Execute()
}
