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

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	file := string(fileBytes)
	file, err = fixDoubleStarProperties(file)
	if err != nil {
		return err
	}

	file, err = fixDoubleBlankProperties(file)
	if err != nil {
		return err
	}

	// TODO(kfcampbell): verify file permission is what we want
	err = os.WriteFile(filename, []byte(file), 0666)
	if err != nil {
		return err
	}

	return nil
}

// fixDoubleStarProperties fixes generated enums where "*" and "Star" are
// both webhook events parsed as "Star"
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

// fixDoubleBlankProperties removes one of the two BLANK enum properties from Repository11.cs
// TODO(kfcampbell): attempt to figure out where this is defined inside GitHub and PR it
func fixDoubleBlankProperties(inputFile string) (string, error) {
	if len(inputFile) < 1 {
		return "", fmt.Errorf("inputFile must not be empty")
	}

	toReplace := `/// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 3,

            /// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 4
`

	// replace duplicate blanks if they exist
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, `/// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 3
		`)
	}
	return inputFile, nil
}
