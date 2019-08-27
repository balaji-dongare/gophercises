package cmd

import (
	"fmt"

	"github.com/balaji-dongare/gophercises/secret/vault"

	"github.com/spf13/cobra"
)

// getCmd is the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get it return api key from secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsFilePath())
		value, err := v.Get(args[0])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return
		}
		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
