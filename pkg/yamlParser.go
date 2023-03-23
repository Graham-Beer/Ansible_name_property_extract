package FileParser

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Step struct{ Name string }

func Parse(file string, listview bool) []string {
	yfile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var data []Step
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	var nameHeader []string
	for _, d := range data {
		var str string

		if d.Name == "" {
			continue
		}
		if listview {
			str = strings.Replace(d.Name, "", "- ", 1)
		}
		if !listview {
			str = d.Name
		}
		nameHeader = append(nameHeader, str)
	}
	return nameHeader
}

func FileExtensionCheck(ext string) bool {
	if filepath.Ext(ext) == ".yml" {
		return true
	}
	if filepath.Ext(ext) == ".yaml" {
		return true
	}
	return false
}
