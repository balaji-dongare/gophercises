package cmd

import (
	"fmt"
	"strings"

	"github.com/balaji-dongare/gophercises/secret/vault"
	"github.com/spf13/cobra"
)

// setCmd  the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set it encrypts and  put's api key into secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsFilePath())
		Plaintext := strings.Join(args[1:], " ")
		err := v.Set(args[0], Plaintext)
		if err != nil {
			fmt.Println("Key not set")
			return
		}
		fmt.Println("Key set")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
