package main

import (
	"bytes"
	"os"
	"flag"
)

type NetworksAndAZs struct {
	outputFile string
}

func (nz NetworksAndAZs) ProcessData() {

	if nz.outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var file *os.File
	if !FileExists(nz.outputFile) {
		file = CreateFile(nz.outputFile)
	}

  defer file.Close()

  var buf bytes.Buffer
  buf.WriteString("network-properties:\n")
  buf.WriteString("  network:\n")
  buf.WriteString("    name:\n")
  buf.WriteString("  service-network:\n")
  buf.WriteString("    name:\n")
  buf.WriteString("  other_availability_zones:\n")
  buf.WriteString("  - name:\n")
  buf.WriteString("  - name:\n")
  buf.WriteString("  singleton_availability_zone:\n")
  buf.WriteString("    name:\n")

	WriteContents(file, buf.String())
}
