package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("exactly one directory name must be provided to run post-processing on")
	}
	dirPath := os.Args[1]
	log.Printf("running Ruby post-processing on directory %v", dirPath)

	files = make(map[string]fs.FileInfo)
	err := filepath.Walk(dirPath, walkFiles)
	if err != nil {
		return fmt.Errorf("could not read directory %v: %v", dirPath, err)
	}

	for path, file := range files {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read file %v: %v", file.Name(), err)
		}
		fileContents := string(fileBytes)
		if len(fileContents) < 1 {
			// skip empty files that are generated in the Ruby SDK
			continue
		}

		fileContents = fixEmptyCaseStatements(fileContents)
		fileContents = fixThumbsUpThumbsDownProperties(fileContents)

		// TODO(kfcampbell): verify file permission is what we want
		err = os.WriteFile(path, []byte(fileContents), 0666)
		if err != nil {
			return err
		}
	}
	// do imports here via cmdline
	// initialize a module and do basic request

	return nil
}

// TODO(kfcampbell) fix this sleazy global
var files map[string]fs.FileInfo

func walkFiles(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error walking files: %v", err)
	}

	// skip directories
	if info.IsDir() {
		return nil
	}
	files[path] = info
	return nil
}

// TODO(kfcampbell): implement and validate
func fixEmptyCaseStatements(inputFile string) string {
	/*
		case mapping_value
		end
	*/

	emptyCaseStatements := regexp.MustCompile(`\s+case mapping_value\n\s+end`)
	inputFile = emptyCaseStatements.ReplaceAllString(inputFile, ``)
	return inputFile
}

// TODO(kfcampbell): implement and validate
func fixThumbsUpThumbsDownProperties(inputFile string) string {
	/*
		One: :One,
		One: :One,
	*/
	thumbsUpThumbsDown := regexp.MustCompile(`(\s+)One: :One,(\n\s+)One: :One,`)
	inputFile = thumbsUpThumbsDown.ReplaceAllString(inputFile, `${1}ThumbsUp: :ThumbsUp,${2}ThumbsDown: :ThumbsDown,`)

	return inputFile
}
