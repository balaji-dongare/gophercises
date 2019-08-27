package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var encodingKey string
var rootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret is secrets manager CLI Application",
}

// Execute it will adds all child commands to the root command.
//It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
}

func secretsFilePath() string {
	dir, _ := os.Getwd()
	secretpath := filepath.Join(dir, ".secrets")
	return secretpath
}
