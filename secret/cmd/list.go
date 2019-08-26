package cmd

import (
	"fmt"

	"github.com/balaji-dongare/gophercises/secret/vault"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is a CLI command to show all key list in secrets ",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsFilePath())
		value, err := v.List()
		if err != nil {
			fmt.Print("\nNo key found")
		}
		if len(value) > 0 {
			for i, key := range value {
				fmt.Printf("%d. %v", i+1, key)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
