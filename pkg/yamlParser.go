package FileParser

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Name string
}

type Step struct {
	Name  string
	Tasks []Task
}

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
	return getSliceData(data, listview)
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

func getSliceData(slice []Step, list bool) []string {
	var nameHeader []string
	for _, s := range slice {
		for _, t := range s.Tasks {
			if t.Name == "" && s.Name == "" {
				continue
			}

			value := t.Name
			if value == "" {
				value = s.Name
			}

			if list {
				value = strings.Replace(value, "", "- ", 1)
			}

			nameHeader = append(nameHeader, value)
		}
	}
	return nameHeader
}
