package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
  "flag"
)

type Errands struct {
  inputFile    string
	outputFile   string
}

func (e Errands) ProcessData() {
  if e.inputFile == "" || e.outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	rawData := GetRaw(e.inputFile)

	var file *os.File
	if !FileExists(e.outputFile) {
		file = CreateFile(e.outputFile)
	}

  defer file.Close()

  // Generic interface to read the file into
	var f interface{}
	err := json.Unmarshal(rawData, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

  // Fetch the top level errands from the json
  m := f.(map[string]interface{})

  var v []interface{}

  // Fetch the errands attribute from the json
	if errandsMap, ok := m["errands"]; ok {
		// Fetch the errands from the errands map
		v = errandsMap.([]interface{})
	} else {
		fmt.Println("Cannot process the input file")
		os.Exit(1)
	}

  s := "errand-config:\n"
  WriteContents(file, s)

	for k := range v {
		node := v[k]

		nodeData := node.(map[string]interface{})
    if nodeData["post_deploy"] == true || nodeData["post_deploy"] == false {
      var buf bytes.Buffer
      s := fmt.Sprintf("  %s: \n", nodeData["name"])
      buf.WriteString(s)

      s = fmt.Sprintf("    %s: %t\n", "post-deploy-state", nodeData["post_deploy"])
      buf.WriteString(s)

      s = fmt.Sprintf("    %s: %s\n", "pre-delete-state", "default")
      buf.WriteString(s)
    	WriteContents(file, buf.String())
    }
  }
}
