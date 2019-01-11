package main

import (
	"flag"
	"os"
)

func main() {
	generate := flag.String("g", "", "properties, resources, errands, network-azs")
	inputFile := flag.String("i", "", "input filename")
	outputFile := flag.String("o", "", "output filename")

	flag.Parse()

	if *generate == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *generate == "properties" {
		p := Properties{inputFile: *inputFile, outputFile: *outputFile}
		p.ProcessData()
	} else if *generate == "resources" {
		r := Resources{inputFile: *inputFile, outputFile: *outputFile}
		r.ProcessData()
	} else if *generate == "errands" {
		e := Errands{inputFile: *inputFile, outputFile: *outputFile}
		e.ProcessData()
	} else if *generate == "network-azs" {
		nz := NetworksAndAZs{outputFile: *outputFile}
		nz.ProcessData()
	}
}
