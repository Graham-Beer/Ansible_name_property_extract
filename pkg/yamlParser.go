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
	var str string
	for _, s := range slice {
		if s.Name == "" {
			continue
		}
		if list {
			str = strings.Replace(s.Name, "", "- ", 1)
		} else {
			str = s.Name
		}
		nameHeader = append(nameHeader, str)
		if len(s.Tasks) > 0 {
			for _, t := range s.Tasks {
				if t.Name == "" {
					continue
				}
				if list {
					str = strings.Replace(t.Name, "", "- ", 1)
				} else {
					str = t.Name
				}
				nameHeader = append(nameHeader, str)
			}
		}
	}
	return nameHeader
}
