package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	log.Printf("running post processing for Go SDK on directory %v", dirPath)

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
			return fmt.Errorf("input file %v must not be empty", file.Name())
		}

		fileContents = fixKiotaFileNameTypeNameError(fileContents, file.Name())

		err = os.WriteFile(path, []byte(fileContents), 0600)
		if err != nil {
			return err
		}
	}

	err = os.Remove(filepath.Join(dirPath, "go.mod"))
	if err != nil {
		return fmt.Errorf("could not remove go.mod file: %v", err)
	}

	err = os.Remove(filepath.Join(dirPath, "go.sum"))
	if err != nil {
		return fmt.Errorf("could not remove go.sum file: %v", err)
	}

	// get the immediate parent directory name as this is the name of the module
	// e.g. if the directory is /path/to/go-sdk, the module name is go-sdk
	_, partialModuleName := filepath.Split(dirPath)
	fullModuleName := fmt.Sprintf("github.com/octokit/%s", partialModuleName)

	cmd := exec.Command("go", "mod", "init", fullModuleName)
	cmd.Dir = dirPath

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	_ = stdout.String()
	stderrOutput := stderr.String()

	if err != nil {
		return fmt.Errorf("could not initialize Go module: %v\nfull error log:\n%s", err, stderrOutput)
	}

	cmd = exec.Command("kiota", "info", "-l", "Go", "--json")
	cmd.Dir = dirPath

	stdout.Reset()
	stderr.Reset()

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("could not run kiota info: %v\nfull error log:\n%s", err, stderr.String())
	}

	// parse the json returned by kiota info, extract the "dependencies" field,
	// and construct a "go get" command for each with the "name" and "version" subfields used.
	// this code could result in exceptions if the format of the command isn't changed.
	// if/when we know for sure that it's stabilized, this can be refactored to use structs
	// and not all of these interface{}s and sleazy casting
	var infoResult map[string]interface{}
	if err := json.Unmarshal(output, &infoResult); err != nil {
		return fmt.Errorf("could not parse kiota info output: %v", err)
	}
	deps := infoResult["dependencies"].([]interface{})
	depsToInstall := make([]string, len(deps))
	for i, d := range deps {
		dep := d.(map[string]interface{})
		name := dep["name"].(string)
		version := dep["version"].(string)

		fullDep := fmt.Sprintf("%s@%s", name, version)
		depsToInstall[i] = fullDep
	}

	for _, dep := range depsToInstall {
		cmd = exec.Command("go", "get", dep)
		cmd.Dir = dirPath

		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("could not get dependencies: %v", err)
		}
		fmt.Printf("installed dependency %s\n", dep)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = dirPath

	stdout.Reset()
	stderr.Reset()

	_, err = cmd.Output()
	if err != nil {
		fmt.Printf("could not run go mod tidy: %v\nfull error log:\n%s", err, stderr.String())
	}

	return nil
}

// TODO(kfcampbell) fix this sleazy global
var files map[string]fs.FileInfo

func walkFiles(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error walking files: %v", err)
	}

	// skip the .git directory
	if strings.Contains(path, ".git") {
		return nil
	}

	// skip other non-dot directories
	if info.IsDir() {
		return nil
	}
	files[path] = info
	return nil
}

func fixKiotaFileNameTypeNameError(inputFile string, fileName string) string {
	if !strings.Contains(fileName, "code_scanning_variant_analysis.go") {
		return inputFile
	}

	// Kiota's generation gives the file name rather than the type name in this
	// specific case, so this workaround fixes it.
	toReplace := "CodeScanningVariantAnalysis_status"
	replaceWith := "CodeScanningVariantAnalysisStatus"

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}
	return inputFile
}
