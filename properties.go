package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Properties struct {
	rawData    []byte
	outputFile string
}

func (p Properties) ProcessData() {
	var file *os.File
	if !FileExists(p.outputFile) {
		file = CreateFile(p.outputFile)
	}

	defer file.Close()

	// Generic interface to read the file into
	var f interface{}
	err1 := json.Unmarshal(p.rawData, &f)
	if err1 != nil {
		fmt.Println("Error parsing JSON: ", err1)
	}

	// Fetch the top level properties from the json
	m := f.(map[string]interface{})

	var v map[string]interface{}

	// Fetch the properties attribute from the json
	if propertiesMap, ok := m["properties"]; ok {
		// Fetch the properties from the properties map
		v = propertiesMap.(map[string]interface{})
	} else {
		fmt.Println("Cannot process the input file")
		os.Exit(1)
	}

	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	s := "product_properties: |\n"
	WriteContents(file, s)

	// To perform the opertion you want
	for _, k := range keys {
		node := v[k]

		nodeData := node.(map[string]interface{})

		if nodeData["configurable"] == true {
			s := fmt.Sprintf("  %s:\n", k)
			WriteContents(file, s)

			if nodeData["type"] == "rsa_cert_credentials" {
				var buf bytes.Buffer
				buf.WriteString("    value: \n")
				buf.WriteString("      private_key_pem: \n")
				buf.WriteString("      cert_pem: \n")
				WriteContents(file, buf.String())
			} else if nodeData["type"] == "secret" {
				var buf bytes.Buffer
				buf.WriteString("    value: \n")
				buf.WriteString("      secret: \n")
				WriteContents(file, buf.String())
			} else if nodeData["type"] == "simple_credentials" {
				var buf bytes.Buffer
				buf.WriteString("    value: \n")
				buf.WriteString("      identity: \n")
				buf.WriteString("      password: \n")
				WriteContents(file, buf.String())
			} else if nodeData["type"] == "collection" {
				var buf bytes.Buffer
				value := nodeData["value"].([]interface{})

				buf.WriteString("    value: \n")
				arrayAdded := false

				for _, item := range value {
					for innerKey, innerVal := range item.(map[string]interface{}) {
						typeAssertedInnerValue := innerVal.(map[string]interface{})
						innerValueType := typeAssertedInnerValue["type"]
						var s string
						if !arrayAdded {
							if innerValueType == "rsa_cert_credentials" {
								s = fmt.Sprintf("    - %s:\n", innerKey)
								buf.WriteString(s)
								buf.WriteString("        private_key_pem: \n")
								buf.WriteString("        cert_pem: \n")
							} else if innerValueType == "secret" {
								s = fmt.Sprintf("    - %s:\n", innerKey)
								buf.WriteString(s)
								buf.WriteString("        secret: \n")
								WriteContents(file, buf.String())
							} else {
								s = fmt.Sprintf("    - %s: %v \n", innerKey, typeAssertedInnerValue["value"])
								buf.WriteString(s)
							}
							arrayAdded = true
						} else {
							if innerValueType == "rsa_cert_credentials" {
								s = fmt.Sprintf("      %s:\n", innerKey)
								buf.WriteString(s)
								buf.WriteString("        private_key_pem: \n")
								buf.WriteString("        cert_pem: \n")
							} else if innerValueType == "secret" {
								s = fmt.Sprintf("      %s:\n", innerKey)
								buf.WriteString(s)
								buf.WriteString("        secret: \n")
								WriteContents(file, buf.String())
							} else {
								s = fmt.Sprintf("      %s: %v \n", innerKey, typeAssertedInnerValue["value"])
								buf.WriteString(s)
							}
						}
					}
					arrayAdded = false
				}
				WriteContents(file, buf.String())
			} else {
				var s string
				value := nodeData["value"]
				if value != nil {
					s = fmt.Sprintf("    value: %v\n", value)
				} else {
					s = fmt.Sprintf("    value: \n")
				}
				WriteContents(file, s)
			}
		}
	}
}
