package FileParser

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Step struct{ Name string }

func Parse(file string) []string {
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
		str := strings.ReplaceAll(d.Name, "{}", "")
		if str == "" {
			continue
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
