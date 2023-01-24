package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("exactly one filename must be provided to run post-processing on")
	}
	filename := os.Args[1]
	log.Printf("running post processing on file %v", filename)

	// open file
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	file := string(fileBytes)
	file, err = fixDoubleStarProperties(file)
	if err != nil {
		return err
	}

	// todo(kfcampbell): verify file permission is what we want
	err = os.WriteFile(filename, []byte(file), 0666)
	if err != nil {
		return err
	}

	return nil
}

// fixDoubleStarProperties returns the input string with
func fixDoubleStarProperties(inputFile string) (string, error) {
	if len(inputFile) < 1 {
		return "", fmt.Errorf("inputFile must not be empty")
	}

	// find instances of "/// Enum Star for value: *" and replace with "Enum All for value: *"
	inputFile = strings.ReplaceAll(inputFile, "/// Enum Star for value: *", "/// Enum All for value: *")

	// find instances of "            Star = 1," and replace with "            All = 1,"
	inputFile = strings.ReplaceAll(inputFile, "            Star = 1,", "            All = 1,")
	return inputFile, nil
}
