package main

import (
	"flag"
)

func main() {
	generate := flag.String("g", "properties", "properties, resources, apply-changes-config")
	inputFile := flag.String("i", "", "input filename")
	outputFile := flag.String("o", "", "output filename")

	flag.Parse()

	rawData := GetRaw(*inputFile)

	if *generate == "properties" {
		p := Properties{rawData: rawData, outputFile: *outputFile}
		p.ProcessData()
	} else if *generate == "resources" {
		r := Resources{rawData: rawData, outputFile: *outputFile}
		r.ProcessData()
	}
}
