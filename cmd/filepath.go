package cmd

import (
	"fmt"

	FileParser "yaml_parser/pkg"

	"github.com/spf13/cobra"
)

var yamlParserCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"yp"},
	Short:   "List yaml name tags",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := FileParser.Parse(args[0])
		for _, n := range name {
			fmt.Println(n)
		}
	},
}

func init() {
	rootCmd.AddCommand(yamlParserCmd)
}
