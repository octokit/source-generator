package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
		fileContents = fixThumbsUpThumbsDownProperties(fileContents)
		fileContents = fixDuplicateValueInEnums(fileContents)
		fileContents = fixClientWorkflowIDErrors(fileContents)

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

func fixClientWorkflowIDErrors(inputFile string) string {
	removeStruct := regexp.MustCompile(`type Components.*ParametersWorkflowIDSchema struct {\n}\n\n\n`)
	inputFile = removeStruct.ReplaceAllString(inputFile, `\n`) // just delete it

	// important! this regex must be run _after_ the above removeStruct regex,
	// otherwise the removeStruct regex will not match anything
	makeString := regexp.MustCompile(`Components.*ParametersWorkflowIDSchema`)
	inputFile = makeString.ReplaceAllString(inputFile, `string`)
	return inputFile
}

func fixDuplicateValueInEnums(inputFile string) string {
	thumbsUp := regexp.MustCompile(`Enum\d+ (Enum(\d+)) = ("\+1")`)
	thumbsDown := regexp.MustCompile(`Enum\d+ (Enum(\d+)) = ("\-1")`)

	inputFile = thumbsUp.ReplaceAllString(inputFile, `Enum${2}ThumbsUp ${1} = ${3}`)
	inputFile = thumbsDown.ReplaceAllString(inputFile, `Enum${2}ThumbsDown ${1} = ${3}`)

	possibleVals := regexp.MustCompile(`Enum(\d{1,3})1,(\n\t+)Enum\d{1,3}1`)
	inputFile = possibleVals.ReplaceAllString(inputFile, `Enum${1}ThumbsUp,${2}Enum${1}ThumbsDown`)

	// important! this regex must be run _after_ the above possibleVals regex,
	// otherwise the possibleVals regex will not match anything.
	singleThumbsUpPossibleValues := regexp.MustCompile(`Enum(\d{1,3})1,(\n\t+)`)
	inputFile = singleThumbsUpPossibleValues.ReplaceAllString(inputFile, `Enum${1}ThumbsUp,${2}`)

	reactionsThumbsUp := regexp.MustCompile(`Enum(\d{1,3})Reactions1 Enum\d{1,3} = ("reactions-\+1")`)
	inputFile = reactionsThumbsUp.ReplaceAllString(inputFile, `Enum${1}ReactionsThumbsUp Enum${1} = ${2}`)

	reactionsThumbsDown := regexp.MustCompile(`Enum(\d{1,3})Reactions1 Enum\d{1,3} = ("reactions--1")`)
	inputFile = reactionsThumbsDown.ReplaceAllString(inputFile, `Enum${1}ReactionsThumbsDown Enum${1} = ${2}`)

	reactionsPossibleVals := regexp.MustCompile(`Enum(\d{1,3})Reactions1,(\n\t+)Enum\d{1,3}Reactions1,`)
	inputFile = reactionsPossibleVals.ReplaceAllString(inputFile, `Enum${1}ReactionsThumbsUp,${2}Enum${1}ReactionsThumbsDown,`)

	postContentThumbs := regexp.MustCompile(`PostContentSchemaContent1 PostContentSchemaContent = ("\+1")(\n\t+)PostContentSchemaContent1 PostContentSchemaContent = ("-1")`)
	inputFile = postContentThumbs.ReplaceAllString(inputFile, `PostContentSchemaContentThumbsUp PostContentSchemaContent = ${1}${2}PostContentSchemaContentThumbsDown PostContentSchemaContent = ${3}`)

	postContentPossibleVals := regexp.MustCompile(`PostContentSchemaContent1,(\n\t+)PostContentSchemaContent1,`)
	inputFile = postContentPossibleVals.ReplaceAllString(inputFile, `PostContentSchemaContentThumbsUp,${1}PostContentSchemaContentThumbsDown,`)

	reactionContentThumbs := regexp.MustCompile(`ReactionContent1 ReactionContent = ("\+1")(\n\t+)ReactionContent1 ReactionContent = ("-1")`)
	inputFile = reactionContentThumbs.ReplaceAllString(inputFile, `ReactionContentThumbsUp ReactionContent = ${1}${2}ReactionContentThumbsDown ReactionContent = ${3}`)

	reactionContentPossibleVals := regexp.MustCompile(`ReactionContent1,(\n\t+)ReactionContent1,`)
	inputFile = reactionContentPossibleVals.ReplaceAllString(inputFile, `ReactionContentThumbsUp,${1}ReactionContentThumbsDown,`)

	return inputFile
}

// TODO(kfcampbell): figure out if some of these can be done with directives instead
func fixThumbsUpThumbsDownProperties(inputFile string) string {
	/*
		// REQUIRED
		1 *int32 `json:"+1,omitempty"`

		// REQUIRED
		1 *int32 `json:"-1,omitempty"`
	*/
	toReplace := `
	// REQUIRED
	1 *int32 ` + "`json:\"+1,omitempty\"`" + `

	// REQUIRED
	1 *int32 ` + "`json:\"-1,omitempty\"`"

	replaceWith := `
	// REQUIRED
	ThumbsUp *int32 ` + "`json:\"+1,omitempty\"`" + `

	// REQUIRED
	ThumbsDown *int32 ` + "`json:\"-1,omitempty\"`"

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		1 *int32 `json:"+1,omitempty"`
		1 *int32 `json:"-1,omitempty"`
	*/
	toReplace = `
	1 *int32 ` + "`json:\"+1,omitempty\"`" + `
	1 *int32 ` + "`json:\"-1,omitempty\"`"

	replaceWith = `
	ThumbsUp *int32 ` + "`json:\"+1,omitempty\"`" + `
	ThumbsDown *int32 ` + "`json:\"-1,omitempty\"`"

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		populate(objectMap, "+1", r.1)
		populate(objectMap, "-1", r.1)
	*/
	toReplace = `
	populate(objectMap, "+1", r.1)
	populate(objectMap, "-1", r.1)
	`

	replaceWith = `
	populate(objectMap, "+1", r.ThumbsUp)
	populate(objectMap, "-1", r.ThumbsDown)
	`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		case "+1":
			err = unpopulate(val, "1", &r.1)
			delete(rawMsg, key)
		case "-1":
			err = unpopulate(val, "1", &r.1)
			delete(rawMsg, key)
	*/

	toReplace = `
		case "+1":
				err = unpopulate(val, "1", &r.1)
				delete(rawMsg, key)
		case "-1":
				err = unpopulate(val, "1", &r.1)
				delete(rawMsg, key)
	`

	replaceWith = `
		case "+1":
			err = unpopulate(val, "+1", &r.ThumbsUp)
			delete(rawMsg, key)
		case "-1":
				err = unpopulate(val, "-1", &r.ThumbsDown)
				delete(rawMsg, key)
	`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		populate(objectMap, "+1", w.1)
		populate(objectMap, "-1", w.1)
	*/
	toReplace = `
	populate(objectMap, "+1", w.1)
	populate(objectMap, "-1", w.1)
	`

	replaceWith = `
	populate(objectMap, "+1", w.ThumbsUp)
	populate(objectMap, "-1", w.ThumbsDown)
	`
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	/*
		case "+1":
				err = unpopulate(val, "1", &w.1)
				delete(rawMsg, key)
		case "-1":
				err = unpopulate(val, "1", &w.1)
				delete(rawMsg, key)
	*/
	toReplace = `
		case "+1":
				err = unpopulate(val, "1", &w.1)
				delete(rawMsg, key)
		case "-1":
				err = unpopulate(val, "1", &w.1)
				delete(rawMsg, key)
	`

	replaceWith = `
		case "+1":
				err = unpopulate(val, "1", &w.ThumbsUp)
				delete(rawMsg, key)
		case "-1":
				err = unpopulate(val, "1", &w.ThumbsDown)
				delete(rawMsg, key)
	`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}
