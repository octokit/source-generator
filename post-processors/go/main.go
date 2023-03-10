package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
		fileContents = fixThumbsUpThumbsDownProperties(fileContents)
		fileContents = fixMissingErrorReferences(fileContents)
		fileContents = dirtyHackToBreakFunctionalityForCompilation(fileContents, file.Name())

		// TODO(kfcampbell): verify file permission is what we want
		err = os.WriteFile(path, []byte(fileContents), 0666)
		if err != nil {
			return err
		}
	}

	// after files are written, initialize a module
	cmd := exec.Command("go", "mod", "init", "github.com/octokit/kiota")
	cmd.Dir = dirPath

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("could not initialize Go module: %v", err)
	}
	log.Printf("output of module initialization: %v", output)

	deps := [6]string{
		"github.com/microsoft/kiota-abstractions-go@v0.17.0",
		"github.com/microsoft/kiota-http-go@v0.14.0",
		"github.com/microsoft/kiota-serialization-form-go@v0.3.0",
		"github.com/microsoft/kiota-serialization-json-go@v0.8.0",
		"github.com/microsoft/kiota-authentication-azure-go@v0.6.0",
		"github.com/microsoft/kiota-serialization-text-go@v0.7.0",
	}

	// run go get on deps
	for _, dep := range deps {
		cmd = exec.Command("go", "get", dep)
		cmd.Dir = dirPath

		output, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("could not get dependencies: %v", err)
		}
	}

	// // run build
	// cmd = exec.Command("go", "build", "./...")
	// cmd.Dir = dirPath

	// output, err = cmd.Output()
	// if err != nil {
	// 	return fmt.Errorf("could not build Go SDK successfully: %v", err)
	// }

	// TODO(kfcampbell): create main.go file for testing
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

func fixImports(inputFile string) string {
	// find: kiota/
	// replace: github.com/octokit/kiota/
	inputFile = strings.ReplaceAll(inputFile, "kiota/", "github.com/octokit/kiota/")
	return inputFile
}

func fixDuplicateValueInEnums(inputFile string) string {
	thumbsUp := regexp.MustCompile(`Enum\d+ (Enum(\d+)) = ("\+1")`)
	thumbsDown := regexp.MustCompile(`Enum\d+ (Enum(\d+)) = ("\-1")`)

	inputFile = thumbsUp.ReplaceAllString(inputFile, `Enum${2}ThumbsUp ${1} = ${3}`)
	inputFile = thumbsDown.ReplaceAllString(inputFile, `Enum${2}ThumbsDown ${1} = ${3}`)

	return inputFile
}

func fixMissingErrorReferences(inputFile string) string {
	toReplace := `CreateStar404ErrorFromDiscriminatorValue`
	replaceWith := `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateMoves403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithCard_403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateProjectCard503ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithProject_403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateProject403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithProject_403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithUsername422ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithProject_403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreatePullRequestMergeResult405ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreatePullRequestMergeResult409ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateWithRepo403ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `CreateMoves503ErrorFromDiscriminatorValue`
	replaceWith = `i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixThumbsUpThumbsDownProperties(inputFile string) string {
	/*
		ONE_REACTIONSPOSTREQUESTBODY_CONTENT ReactionsPostRequestBody_content = iota
		ONE_REACTIONSPOSTREQUESTBODY_CONTENT
	*/
	toReplace := `ONE_REACTIONSPOSTREQUESTBODY_CONTENT ReactionsPostRequestBody_content = iota
    ONE_REACTIONSPOSTREQUESTBODY_CONTENT`

	replaceWith := `THUMBSUP_REACTIONSPOSTREQUESTBODY_CONTENT ReactionsPostRequestBody_content = iota
    THUMBSDOWN_REACTIONSPOSTREQUESTBODY_CONTENT`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
				case "One":
		            result = ONE_REACTIONSPOSTREQUESTBODY_CONTENT
		        case "One":
		            result = ONE_REACTIONSPOSTREQUESTBODY_CONTENT
	*/
	toReplace = `result := ONE_REACTIONSPOSTREQUESTBODY_CONTENT
    switch v {
        case "One":
            result = ONE_REACTIONSPOSTREQUESTBODY_CONTENT
        case "One":
            result = ONE_REACTIONSPOSTREQUESTBODY_CONTENT`

	replaceWith = `result := THUMBSUP_REACTIONSPOSTREQUESTBODY_CONTENT
    switch v {
        case "+1":
            result = THUMBSUP_REACTIONSPOSTREQUESTBODY_CONTENT
        case "-1":
            result = THUMBSDOWN_REACTIONSPOSTREQUESTBODY_CONTENT`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		ONE_REACTION_CONTENT Reaction_content = iota
		ONE_REACTION_CONTENT
	*/

	toReplace = `ONE_REACTION_CONTENT Reaction_content = iota
    ONE_REACTION_CONTENT`

	replaceWith = `THUMBSUP_REACTION_CONTENT Reaction_content = iota
    THUMBSDOWN_REACTION_CONTENT`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		result := ONE_REACTION_CONTENT
		switch v {
			case "One":
				result = ONE_REACTION_CONTENT
			case "One":
				result = ONE_REACTION_CONTENT
	*/

	toReplace = `result := ONE_REACTION_CONTENT
    switch v {
        case "One":
            result = ONE_REACTION_CONTENT
        case "One":
            result = ONE_REACTION_CONTENT`

	replaceWith = `result := THUMBSUP_REACTION_CONTENT
    switch v {
        case "+1":
            result = THUMBSUP_REACTION_CONTENT
        case "-1":
            result = THUMBSDOWN_REACTION_CONTENT`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		result := ONE_REACTION_CONTENT
		switch v {
			case "One":
				result = ONE_REACTION_CONTENT
			case "One":
				result = ONE_REACTION_CONTENT
	*/

	toReplace = `result := ONE_REACTION_CONTENT
    switch v {
        case "One":
            result = ONE_REACTION_CONTENT
        case "One":
            result = ONE_REACTION_CONTENT`

	replaceWith = `result := THUMBSUP_REACTION_CONTENT
    switch v {
        case "+1":
            result = THUMBSUP_REACTION_CONTENT
        case "-1":
            result = THUMBSDOWN_REACTION_CONTENT`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func dirtyHackToBreakFunctionalityForCompilation(inputFile string, filename string) string {
	toReplace := `for i, v := range res {
        val[i] = *(v.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly))
    }`
	replaceWith := `for i, _ := range res {
		//val[i] = *(v.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly))
		t := serialization.DateOnly{}
		val[i] = t
	}`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateDateOnlyFromDiscriminatorValue, errorMapping)`
	replaceWith = `res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateBasicErrorFromDiscriminatorValue, errorMapping)`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	// change imports if we're in versions_request_builder.go
	if strings.Contains(filename, "versions_request_builder.go") {
		toReplace = `import (
    "context"
    i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830 "github.com/octokit/kiota/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)`

		replaceWith = `import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830 "github.com/octokit/kiota/models"
)`

		if strings.Contains(inputFile, toReplace) {
			inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
		}
	}

	toReplace = `// ItemStarredRepositoryable
type ItemStarredRepositoryable interface {
    IAdditionalDataHolder
}
`

	replaceWith = `
import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemStarredRepositoryable
type ItemStarredRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `m.SetAdditionalData(make(map[string]any))
    is_alphanumericValue := "true"
    m.SetIsAlphanumeric(&is_alphanumericValue)`

	replaceWith = `m.SetAdditionalData(make(map[string]any))
	is_alphanumericValue := true
	m.SetIsAlphanumeric(&is_alphanumericValue)`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `requestConfiguration *ApiClientGetRequestConfiguration`
	replaceWith = `requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}
