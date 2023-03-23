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
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		fileExtension := FileParser.FileExtensionCheck(file)
		if !fileExtension {
			fmt.Println("Incorrect file type! Please parse '.yml' or '.yaml' extensions only.")
			return
		}
		name := FileParser.Parse(file)
		for _, n := range name {
			fmt.Println(n)
		}
	},
}

func init() {
	rootCmd.AddCommand(yamlParserCmd)
	yamlParserCmd.PersistentFlags().StringP("file", "f", "", "Add yaml file path. Must be in the format of '.yml' or '.yaml'")
	if err := yamlParserCmd.MarkPersistentFlagRequired("file"); err != nil {
		fmt.Println(err.Error())
		return
	}
}
