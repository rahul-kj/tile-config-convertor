package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func FileExists(fileName string) bool {
	var fileExists bool
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			fileExists = false
		} else {
			fileExists = true
		}
	}
	os.Remove(fileName)
	return fileExists
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteContents(file *os.File, contents string) {
	w := bufio.NewWriter(file)
	_, err := w.WriteString(contents)
	if err != nil {
		fmt.Println("Error writing to the file: ", err)
	}
	w.Flush()
}

func GetRaw(file string) []byte {

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}
