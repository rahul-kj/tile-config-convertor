package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Resources struct {
	rawData    []byte
	outputFile string
}

func (r Resources) ProcessData() {
	var file *os.File
	if !FileExists(r.outputFile) {
		file = CreateFile(r.outputFile)
	}

	defer file.Close()

	// Generic interface to read the file into
	var f interface{}
	err := json.Unmarshal(r.rawData, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// Fetch the top level properties from the json
	m := f.(map[string]interface{})

	var resources []interface{}

	// Fetch the resources attribute from the json
	if resourcesMap, ok := m["resources"]; ok {
		// Fetch the properties from the properties map
		resources = resourcesMap.([]interface{})
	} else {
		fmt.Println("Cannot process the input file")
		os.Exit(1)
	}

	s := "resource-config:\n"
	WriteContents(file, s)

	for _, item := range resources {
		value := item.(map[string]interface{})
		if int(value["instances_best_fit"].(float64)) != 0 {
			s := fmt.Sprintf("  %v:\n", value["identifier"])
			WriteContents(file, s)

			s = fmt.Sprintf("    instances: %v\n", value["instances_best_fit"])
			WriteContents(file, s)

			s = fmt.Sprintf("    instance_type:\n      id: %v\n", value["instance_type_best_fit"])
			WriteContents(file, s)

			if _, ok := value["persistent_disk_mb"]; ok {
				s := fmt.Sprintf("    persistent_disk:\n      size_mb: \"%v\"\n", value["persistent_disk_best_fit"])
				WriteContents(file, s)
			}
		}
	}
}
