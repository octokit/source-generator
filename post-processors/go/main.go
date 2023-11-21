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

		fileContents = fixImports(fileContents)
		fileContents = fixCreateDateOnlyFromDiscriminatorValue(fileContents, file.Name())
		fileContents = dirtyHackForVersionsRequestBuilder(fileContents, file.Name())
		fileContents = fixKiotaNonDeterminism(fileContents, file.Name())

		// TODO(kfcampbell): verify file permission is what we want
		err = os.WriteFile(path, []byte(fileContents), 0666)
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

	cmd := exec.Command("go", "mod", "init", "github.com/octokit/go-sdk")
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

	output, err = cmd.Output()
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

// these fixes are working around bugs or limitations in Kiota and/or our schema
func fixImports(inputFile string) string {
	// find: go-sdk/
	// replace: github.com/octokit/go-sdk/
	inputFile = strings.ReplaceAll(inputFile, `"octokit/`, `"github.com/octokit/go-sdk/github/octokit/`)
	return inputFile
}

func dirtyHackForVersionsRequestBuilder(inputFile string, fileName string) string {
	if !strings.Contains(fileName, "versions_request_builder.go") {
		return inputFile
	}

	toReplace := `
	"github.com/microsoft/kiota-abstractions-go/serialization"`
	replaceWith := ``

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `for i, v := range res {
        if v != nil {
            val[i] = *(v.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly))
        }
    }`
	replaceWith = `// for i, v := range res {
    //     if v != nil {
    //         val[i] = *(v.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly))
    //     }`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixCreateDateOnlyFromDiscriminatorValue(inputFile string, filename string) string {
	toReplace := `res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateDateOnlyFromDiscriminatorValue, errorMapping)`
	replaceWith := `res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue, errorMapping)`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateDateOnlyFromDiscriminatorValue, errorMapping)`
	replaceWith = `res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateKeySimpleFromDiscriminatorValue, errorMapping)`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixKiotaNonDeterminism(inputFile string, fileName string) string {
	if !strings.Contains(fileName, "item_starred_repository.go") {
		return inputFile
	}
	// the extra space at the top of this string is because the Kiota
	// generation leaves it in, but saving this file causes the Go formatter
	// to remove the space and then the string won't match
	toReplace := `// ItemStarredRepositoryable` + " " + `
type ItemStarredRepositoryable interface {
    IAdditionalDataHolder
}`
	replaceWith := `// ItemStarredRepositoryable
type ItemStarredRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}
	return inputFile
}
