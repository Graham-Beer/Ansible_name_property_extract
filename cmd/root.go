package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.4"

var rootCmd = &cobra.Command{
	Use:     "YamlNameParser",
	Version: version,
	Short:   "A Parser to use with Yaml files to get the name property",
	Long:    `A Parser to use with an Ansible file to pull out the name property details.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
