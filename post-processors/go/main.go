package main

import (
	"bytes"
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
		fileContents = fixPackageNameInAPIClient(fileContents, file.Name())
		fileContents = fixDuplicateStruct(fileContents, file.Name())
		fileContents = removeModelsKiotaDoesNotCleanUp(fileContents)
		fileContents = dirtyHackForVersionsRequestBuilder(fileContents, file.Name())

		// TODO(kfcampbell): verify file permission is what we want
		err = os.WriteFile(path, []byte(fileContents), 0666)
		if err != nil {
			return err
		}
	}

	// after files are written, initialize a module
	cmd := exec.Command("go", "mod", "init", "github.com/octokit/kiota")
	cmd.Dir = dirPath

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	stdoutOutput := stdout.String()
	stderrOutput := stderr.String()

	if err != nil {
		return fmt.Errorf("could not initialize Go module: %v\nfull error log:\n%s", err, stderrOutput)
	}

	log.Printf("output of module initialization: %v", stdoutOutput)

	deps := [7]string{
		"github.com/microsoft/kiota-abstractions-go@v1.3.0",
		"github.com/microsoft/kiota-http-go@v1.1.0",
		"github.com/microsoft/kiota-serialization-form-go@v1.0.0",
		"github.com/microsoft/kiota-serialization-json-go@v1.0.4",
		"github.com/microsoft/kiota-authentication-azure-go@v1.0.1",
		"github.com/microsoft/kiota-serialization-text-go@v1.0.0",
		"github.com/microsoft/kiota-serialization-multipart-go@v1.0.0",
	}

	// run go get on deps
	for _, dep := range deps {
		cmd = exec.Command("go", "get", dep)
		cmd.Dir = dirPath

		_, err := cmd.Output()
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

// these fixes are working around bugs or limitations in Kiota and/or our schema
func fixImports(inputFile string) string {
	// find: kiota/
	// replace: github.com/octokit/kiota/
	inputFile = strings.ReplaceAll(inputFile, "kiota/", "github.com/octokit/kiota/")
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

func removeModelsKiotaDoesNotCleanUp(inputFile string) string {
	toReplace := `// If specified, only code scanning alerts with this severity will be returned.
    SeverityAsCodeScanningAlertSeverity *i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CodeScanningAlertSeverity ` + "`uriparametername:\"severity\"`"
	replaceWith := ``

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `// If specified, only code scanning alerts with this state will be returned.
    StateAsCodeScanningAlertStateQuery *i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CodeScanningAlertStateQuery ` + "`uriparametername:\"state\"`"

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixDuplicateStruct(inputFile string, fileName string) string {
	if !strings.Contains(fileName, "import_escapedable.go") {
		return inputFile
	}

	toReplace := `// ImportEscapedable
type ImportEscapedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorsCount()(*int32)
    GetAuthorsUrl()(*string)
    GetCommitCount()(*int32)
    GetErrorMessage()(*string)
    GetFailedStep()(*string)
    GetHasLargeFiles()(*bool)
    GetHtmlUrl()(*string)
    GetImportPercent()(*int32)
    GetLargeFilesCount()(*int32)
    GetLargeFilesSize()(*int32)
    GetMessage()(*string)
    GetProjectChoices()([]Import_project_choicesable)
    GetPushPercent()(*int32)
    GetRepositoryUrl()(*string)
    GetStatus()(*Import_status)
    GetStatusText()(*string)
    GetSvcRoot()(*string)
    GetSvnRoot()(*string)
    GetTfvcProject()(*string)
    GetUrl()(*string)
    GetUseLfs()(*bool)
    GetVcs()(*string)
    GetVcsUrl()(*string)
    SetAuthorsCount(value *int32)()
    SetAuthorsUrl(value *string)()
    SetCommitCount(value *int32)()
    SetErrorMessage(value *string)()
    SetFailedStep(value *string)()
    SetHasLargeFiles(value *bool)()
    SetHtmlUrl(value *string)()
    SetImportPercent(value *int32)()
    SetLargeFilesCount(value *int32)()
    SetLargeFilesSize(value *int32)()
    SetMessage(value *string)()
    SetProjectChoices(value []Import_project_choicesable)()
    SetPushPercent(value *int32)()
    SetRepositoryUrl(value *string)()
    SetStatus(value *Import_status)()
    SetStatusText(value *string)()
    SetSvcRoot(value *string)()
    SetSvnRoot(value *string)()
    SetTfvcProject(value *string)()
    SetUrl(value *string)()
    SetUseLfs(value *bool)()
    SetVcs(value *string)()
    SetVcsUrl(value *string)()
}`

	replaceWith := ``

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
	replaceWith = `res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i158396662f32fe591e8faa247af18558546841dba91f24f5c824e11e34188830.CreateKeySimpleFromDiscriminatorValue, errorMapping)`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixPackageNameInAPIClient(inputFile string, filename string) string {
	if strings.Contains(filename, "api_client.go") {
		toReplace := `package kiota`
		replaceWith := `package main`

		if strings.Contains(inputFile, toReplace) {
			inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
		}
	}
	return inputFile
}
