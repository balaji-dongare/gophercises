package cmd

import (
	"fmt"

	"github.com/balaji-dongare/gophercises/secret/vault"
	"github.com/spf13/cobra"
)

// setCmd  the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set it put's api key into secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsFilePath())
		err := v.Set(args[0], args[1])
		if err != nil {
			fmt.Println("Key not set")
		}
		fmt.Println("Key set")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
