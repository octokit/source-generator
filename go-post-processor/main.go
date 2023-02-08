package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
	log.Printf("running post processing on directory %v", dirPath)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("could not read directory %v: %v", dirPath, err)
	}

	// make output directory to avoid processing in-place until this is stable
	// TODO(kfcampbell): verify directory permission is what we want
	outputDirName := "output"
	err = os.Mkdir(outputDirName, 0755)
	if err != nil {
		return fmt.Errorf("error creating %v directory: %v", outputDirName, err)
	}
	packageName := "gh"
	err = os.Mkdir(fmt.Sprintf("%v/%v", outputDirName, packageName), 0755)
	if err != nil {
		return fmt.Errorf("error creating %v directory: %v", packageName, err)
	}

	for _, file := range files {
		fileBytes, err := os.ReadFile(fmt.Sprintf("%v/%v", dirPath, file.Name()))
		if err != nil {
			return fmt.Errorf("could not read file %v: %v", file.Name(), err)
		}
		fileContents := string(fileBytes)
		if len(fileContents) < 1 {
			return fmt.Errorf("input file %v must not be empty", file.Name())
		}

		fileContents = fixPackageNaming(fileContents, packageName)

		// TODO(kfcampbell): verify file permission is what we want
		err = os.WriteFile(fmt.Sprintf("%v/%v/%v", outputDirName, packageName, file.Name()), []byte(fileContents), 0666)
		if err != nil {
			return err
		}
	}

	// after files are written, initialize a module
	cmd := exec.Command("go", "mod", "init", "github.com/nickfloyd/source-generator/go-post-processor/output")
	cmd.Dir = outputDirName

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("could not initialize Go module: %v", err)
	}
	log.Printf("output of module initialization: %v", output)

	// run go get on deps
	cmd = exec.Command("go", "get", "github.com/Azure/azure-sdk-for-go/sdk/azcore")
	cmd.Dir = outputDirName

	output, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("could not get dependencies: %v", err)
	}

	// run build
	cmd = exec.Command("go", "build", "./...")
	cmd.Dir = outputDirName

	output, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("could not build Go SDK successfully: %v", err)
	}

	// TODO(kfcampbell): create main.go file for testing
	return nil
}

func fixPackageNaming(inputFile string, packageName string) string {
	toReplace := "package go"
	replaceWith := fmt.Sprintf("package %v", packageName)
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}
	return inputFile
}

func fixThumbsUpThumbsDownProperties(inputFile string) (string, error) {

	// conditions for needing to fix thumbs up/thumbs down issue:
	// filename contains "Reaction"

	toReplace := ``
	replaceWith := ``

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}
	return inputFile, nil
}
